package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/irfanputra/devpod-playground/pkg/interfaces"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
)

func TestNewDependencies(t *testing.T) {
	tests := []struct {
		name      string
		mysqlConn string
		wantErr   bool
	}{
		{
			name:      "valid connection",
			mysqlConn: "user:pass@tcp(localhost:3306)/db",
			wantErr:   false,
		},
		{
			name:      "invalid connection",
			mysqlConn: "invalid",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewDependencies(tt.mysqlConn)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDependencies() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

type mockDB struct {
	pingErr     error
	CloseCalled bool
}

var _ interfaces.DB = (*mockDB)(nil)

func (m *mockDB) PingContext(ctx context.Context) error {
	return m.pingErr
}

func (m *mockDB) Close() error {
	m.CloseCalled = true
	return nil
}

type mockRedis struct {
	pingErr     error
	CloseCalled bool
}

var _ interfaces.Redis = (*mockRedis)(nil)

func (m *mockRedis) Ping(ctx context.Context) *redis.StatusCmd {
	cmd := redis.NewStatusCmd(ctx, "PING")
	if m.pingErr != nil {
		cmd.SetErr(m.pingErr)
	}
	return cmd
}

func (m *mockRedis) Close() error {
	m.CloseCalled = true
	return nil
}

type mockKafkaWriter struct {
	writeErr    error
	CloseCalled bool
}

var _ interfaces.KafkaWriter = (*mockKafkaWriter)(nil)

func (m *mockKafkaWriter) WriteMessages(ctx context.Context, msgs ...kafka.Message) error {
	return m.writeErr
}

func (m *mockKafkaWriter) Close() error {
	m.CloseCalled = true
	return nil
}

func TestDependencies_Check(t *testing.T) {
	tests := []struct {
		name          string
		dbPingErr     error
		redisPingErr  error
		kafkaWriteErr error
		wantErr       bool
	}{
		{
			name:    "all services healthy",
			wantErr: false,
		},
		{
			name:      "mysql failure",
			dbPingErr: fmt.Errorf("mysql error"),
			wantErr:   true,
		},
		{
			name:         "redis failure",
			redisPingErr: fmt.Errorf("redis error"),
			wantErr:      true,
		},
		{
			name:          "kafka failure",
			kafkaWriteErr: fmt.Errorf("kafka error"),
			wantErr:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deps := &Dependencies{
				db:     &mockDB{pingErr: tt.dbPingErr},
				rdb:    &mockRedis{pingErr: tt.redisPingErr},
				writer: &mockKafkaWriter{writeErr: tt.kafkaWriteErr},
			}

			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()

			err := deps.Check(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Dependencies.Check() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDependencies_Close(t *testing.T) {
	db := &mockDB{}
	rdb := &mockRedis{}
	writer := &mockKafkaWriter{}

	deps := &Dependencies{
		db:     db,
		rdb:    rdb,
		writer: writer,
	}

	deps.Close()

	// Verify that Close was called on all dependencies
	if db.CloseCalled != true {
		t.Error("expected db.Close() to be called")
	}
	if rdb.CloseCalled != true {
		t.Error("expected rdb.Close() to be called")
	}
	if writer.CloseCalled != true {
		t.Error("expected writer.Close() to be called")
	}
}
