package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// MySQL Connection
	mysqlConn := fmt.Sprintf("%s:%s@tcp(mysql:3306)/%s",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_DATABASE"))

	db, err := sql.Open("mysql", mysqlConn)
	if err != nil {
		log.Fatal("Failed to connect to MySQL:", err)
	}
	defer db.Close()

	// Test MySQL connection
	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping MySQL:", err)
	}
	fmt.Println("Successfully connected to MySQL")

	// Redis Connection
	rdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
	defer rdb.Close()

	// Test Redis connection
	ctx := context.Background()
	if err := rdb.Set(ctx, "test_key", "test_value", time.Hour).Err(); err != nil {
		log.Fatal("Failed to set Redis key:", err)
	}
	fmt.Println("Successfully connected to Redis")

	// Kafka Connection
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"kafka:29092"},
		Topic:   "test-topic",
	})
	defer writer.Close()

	// Test Kafka connection
	msg := kafka.Message{
		Value: []byte("test message"),
	}
	if err := writer.WriteMessages(ctx, msg); err != nil {
		log.Fatal("Failed to write to Kafka:", err)
	}
	fmt.Println("Successfully connected to Kafka")

	fmt.Println("All services are connected and working!")
}
