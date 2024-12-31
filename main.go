package main

import (
	"context"
	"database/sql"
	"fmt"

	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/irfanputra/devpod-playground/pkg/interfaces"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
)

// Dependencies represents the dependencies for the application.
type Dependencies struct {
	db     interfaces.DB
	rdb    interfaces.Redis
	writer interfaces.KafkaWriter
}

// NewDependencies creates a new instance of Dependencies.
func NewDependencies(mysqlConn string) (*Dependencies, error) {
	db, err := sql.Open("mysql", mysqlConn)
	if err != nil {
		return nil, err
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"kafka:29092"},
		Topic:   "topic1",
	})

	return &Dependencies{
		db:     db,
		rdb:    rdb,
		writer: writer,
	}, nil
}

// Check checks the health of the dependencies.
func (d *Dependencies) Check(ctx context.Context) error {
	// Check MySQL
	if err := d.db.PingContext(ctx); err != nil {
		return fmt.Errorf("mysql check failed: %v", err)
	}

	// Check Redis
	if _, err := d.rdb.Ping(ctx).Result(); err != nil {
		return fmt.Errorf("redis check failed: %v", err)
	}

	// Check Kafka by attempting to write a test message
	msg := kafka.Message{
		Value: []byte("health check"),
	}
	if err := d.writer.WriteMessages(ctx, msg); err != nil {
		return fmt.Errorf("kafka check failed: %v", err)
	}

	return nil
}

func (d *Dependencies) Close() {
	d.db.Close()
	d.rdb.Close()
	d.writer.Close()
}

var deps *Dependencies

func livenessHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if err := deps.Check(ctx); err != nil {
		log.Printf("Health check failed: %v", err)
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, "Service Unavailable: %v", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")
}

func init() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	mysqlConn := fmt.Sprintf("%s:%s@tcp(mysql:3306)/%s",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_DATABASE"))

	deps, err := NewDependencies(mysqlConn)
	if err != nil {
		log.Fatal(err)
	}
	defer deps.Close()

	http.HandleFunc("/liveness", livenessHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Println("Starting HTTP server on :8080...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("HTTP server error:", err)
		}
	}()

	<-stop
	log.Println("\nShutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Println("Server shutdown error:", err)
	}

	deps.Close()

	log.Println("Server stopped")
}
