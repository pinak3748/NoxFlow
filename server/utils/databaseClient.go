package utils

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type DatabaseClient struct {
	db          *sql.DB
	logBatch    []*LogData
	usageBatch  []*UsageData
	batchSize   int
	batchTicker *time.Ticker
}

type LogData struct {
	Timestamp     time.Time
	ContainerName string
	LogMessage    string
}

type UsageData struct {
	Timestamp     time.Time
	ContainerID   string
	CPUPercent    float64
	MemoryPercent float64
}

// NewDatabaseClient creates a new database client with batch processing

func NewDatabaseClient(connStr string, batchSize int, flushInterval time.Duration) (*DatabaseClient, error) {
	// Validate inputs
	if batchSize <= 0 {
		return nil, errors.New("batch size must be positive")
	}
	if flushInterval <= 0 {
		return nil, errors.New("flush interval must be positive")
	}

	// Connect to database with retry logic
	var dbInstance *sql.DB
	var err error
	for retries := 5; retries > 0; retries-- {
		dbInstance, err = sql.Open("postgres", connStr)
		if err != nil {
			if retries == 1 {
				return nil, fmt.Errorf("failed to connect to database after retries: %v", err)
			}
			time.Sleep(2 * time.Second)
			continue
		}

		// Test connection
		if err := dbInstance.Ping(); err != nil {
			if retries == 1 {
				return nil, fmt.Errorf("failed to ping database after retries: %v", err)
			}
			time.Sleep(2 * time.Second)
			continue
		}
		break
	}

	// Configure connection pool
	dbInstance.SetMaxOpenConns(25)
	dbInstance.SetMaxIdleConns(5)
	dbInstance.SetConnMaxLifetime(5 * time.Minute)

	// // Run migrations
	// if err := runMigrations(dbInstance); err != nil {
	// 	return nil, fmt.Errorf("failed to run migrations: %v", err)
	// }

	client := &DatabaseClient{
		db:          dbInstance,
		batchSize:   batchSize,
		logBatch:    make([]*LogData, 0, batchSize),
		usageBatch:  make([]*UsageData, 0, batchSize),
		batchTicker: time.NewTicker(flushInterval),
	}

	go client.periodicFlush()

	return client, nil
}

func runMigrations(db *sql.DB) error {
	// First, ensure the schema_migrations table exists
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS schema_migrations (
            version bigint NOT NULL,
            dirty boolean NOT NULL,
            CONSTRAINT schema_migrations_pkey PRIMARY KEY (version)
        );
    `)
	if err != nil {
		return fmt.Errorf("failed to create schema_migrations table: %v", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to create migration driver: %v", err)
	}

	migrationsPath := "file://db/migrations"
	m, err := migrate.NewWithDatabaseInstance(
		migrationsPath,
		"postgres",
		driver,
	)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %v", err)
	}

	// Force the version to 0 if there's an error
	if err := m.Force(0); err != nil {
		return fmt.Errorf("failed to force migration version: %v", err)
	}

	// Run migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to run migrations: %v", err)
	}

	return nil
}

// AddLog adds a log entry to the batch
func (c *DatabaseClient) AddLog(log *LogData) error {
	c.logBatch = append(c.logBatch, log)

	if len(c.logBatch) >= c.batchSize {
		return c.FlushLogs()
	}
	return nil
}

// AddUsage adds a usage statistics entry to the batch
func (c *DatabaseClient) AddUsage(usage *UsageData) error {
	c.usageBatch = append(c.usageBatch, usage)

	if len(c.usageBatch) >= c.batchSize {
		return c.FlushUsage()
	}
	return nil
}

// FlushLogs writes all batched logs to the database
func (c *DatabaseClient) FlushLogs() error {
	if len(c.logBatch) == 0 {
		return nil
	}

	tx, err := c.db.BeginTx(context.Background(), nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}

	stmt, err := tx.Prepare(`
		INSERT INTO container_logs (timestamp, container_name, log_message)
		VALUES ($1, $2, $3)
	`)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to prepare statement: %v", err)
	}

	for _, log := range c.logBatch {
		_, err := stmt.Exec(log.Timestamp, log.ContainerName, log.LogMessage)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to insert log: %v", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	c.logBatch = c.logBatch[:0]
	return nil
}

// FlushUsage writes all batched usage statistics to the database
func (c *DatabaseClient) FlushUsage() error {
	if len(c.usageBatch) == 0 {
		return nil
	}

	tx, err := c.db.BeginTx(context.Background(), nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}

	stmt, err := tx.Prepare(`
		INSERT INTO container_usage (timestamp, container_id, cpu_percent, memory_percent)
		VALUES ($1, $2, $3, $4)
	`)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to prepare statement: %v", err)
	}

	for _, usage := range c.usageBatch {
		_, err := stmt.Exec(usage.Timestamp, usage.ContainerID, usage.CPUPercent, usage.MemoryPercent)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to insert usage stats: %v", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	c.usageBatch = c.usageBatch[:0]
	return nil
}

// periodicFlush runs in the background and flushes batches periodically
func (c *DatabaseClient) periodicFlush() {
	for range c.batchTicker.C {
		if err := c.FlushLogs(); err != nil {
			fmt.Printf("Error flushing logs: %v\n", err)
		}
		if err := c.FlushUsage(); err != nil {
			fmt.Printf("Error flushing usage stats: %v\n", err)
		}
	}
}

// Close closes the database connection and flushes any remaining batches
func (c *DatabaseClient) Close() error {
	c.batchTicker.Stop()

	if err := c.FlushLogs(); err != nil {
		return err
	}
	if err := c.FlushUsage(); err != nil {
		return err
	}

	return c.db.Close()
}
