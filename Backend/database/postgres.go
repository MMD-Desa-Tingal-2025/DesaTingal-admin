// backend/database/postgres.go
package database

import (
    "database/sql"
    "fmt"
    "log"
    "time"

    _ "github.com/lib/pq"

    "path/filepath"
    "github.com/golang-migrate/migrate/v4"
    "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/golang-migrate/migrate/v4/source/file"
)

type PostgresDB struct {
    DB *sql.DB
}

// NewPostgresDB creates a new PostgreSQL database connection
func NewPostgresDB(databaseURL string) (*PostgresDB, error) {
    // Open database connection
    db, err := sql.Open("postgres", databaseURL)
    if err != nil {
        return nil, fmt.Errorf("failed to open database: %w", err)
    }

    // Configure connection pool
    db.SetMaxOpenConns(25)                 // Maximum number of open connections
    db.SetMaxIdleConns(25)                 // Maximum number of idle connections
    db.SetConnMaxLifetime(5 * time.Minute) // Maximum connection lifetime
    db.SetConnMaxIdleTime(time.Minute)     // Maximum connection idle time

    // Test the connection
    if err = db.Ping(); err != nil {
        db.Close()
        return nil, fmt.Errorf("failed to ping database: %w", err)
    }

    postgresDB := &PostgresDB{DB: db}
    
    // Run initial setup
    if err := postgresDB.setup(); err != nil {
        db.Close()
        return nil, fmt.Errorf("failed to setup database: %w", err)
    }

    log.Println("Successfully connected to PostgreSQL database")
    return postgresDB, nil
}

// setup runs initial database setup
func (p *PostgresDB) setup() error {
    // Create any initial tables or run migrations here
    // For now, just test the connection
    return p.healthCheck()
}

// Close closes the database connection
func (p *PostgresDB) Close() error {
    if p.DB != nil {
        log.Println("Closing database connection...")
        return p.DB.Close()
    }
    return nil
}

// healthCheck verifies the database connection is healthy
func (p *PostgresDB) healthCheck() error {
    if err := p.DB.Ping(); err != nil {
        return fmt.Errorf("database health check failed: %w", err)
    }
    return nil
}

// GetDB returns the underlying sql.DB instance
func (p *PostgresDB) GetDB() *sql.DB {
    return p.DB
}

// ExecuteQuery executes a SQL query and returns the result
func (p *PostgresDB) ExecuteQuery(query string, args ...interface{}) (*sql.Rows, error) {
    rows, err := p.DB.Query(query, args...)
    if err != nil {
        return nil, fmt.Errorf("failed to execute query: %w", err)
    }
    return rows, nil
}

// ExecuteQueryRow executes a SQL query that returns a single row
func (p *PostgresDB) ExecuteQueryRow(query string, args ...interface{}) *sql.Row {
    return p.DB.QueryRow(query, args...)
}

// ExecuteExec executes a SQL statement (INSERT, UPDATE, DELETE)
func (p *PostgresDB) ExecuteExec(query string, args ...interface{}) (sql.Result, error) {
    result, err := p.DB.Exec(query, args...)
    if err != nil {
        return nil, fmt.Errorf("failed to execute statement: %w", err)
    }
    return result, nil
}

// CreateTable creates a table if it doesn't exist
func (p *PostgresDB) CreateTable(tableName, schema string) error {
    query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s)", tableName, schema)
    _, err := p.DB.Exec(query)
    if err != nil {
        return fmt.Errorf("failed to create table %s: %w", tableName, err)
    }
    log.Printf("Table %s created/verified successfully", tableName)
    return nil
}

// RunMigration runs a simple migration
func (p *PostgresDB) RunMigration() error {
    // Example migration - create a simple logs table
    schema := `
        id SERIAL PRIMARY KEY,
        message TEXT NOT NULL,
        level VARCHAR(10) DEFAULT 'info',
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    `
    
    if err := p.CreateTable("application_logs", schema); err != nil {
        return err
    }
    
    // Add indexes
    indexQuery := `CREATE INDEX IF NOT EXISTS idx_logs_created_at ON application_logs(created_at)`
    if _, err := p.DB.Exec(indexQuery); err != nil {
        return fmt.Errorf("failed to create index: %w", err)
    }
    
    log.Println("Database migration completed successfully")
    return nil
}

func (p *PostgresDB) RunSQLMigration() error {
	driver, err := postgres.WithInstance(p.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to create migration driver: %w", err)
	}

	// Lokasi folder migrations
	path, err := filepath.Abs("./migrations")
	if err != nil {
		return fmt.Errorf("failed to resolve migration path: %w", err)
	}
    // Ganti backslash dengan slash untuk kompatibilitas
    fixedPath := filepath.ToSlash(path)

	m, err := migrate.NewWithDatabaseInstance(
		"file://" + fixedPath,
		"postgres", driver,
	)
	if err != nil {
		return fmt.Errorf("failed to initialize migration: %w", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migration failed: %w", err)
	}

	log.Println("SQL migrasi berhasil dijalankan.")
	return nil
}

// LogMessage logs a message to the database
func (p *PostgresDB) LogMessage(message, level string) error {
    query := `INSERT INTO application_logs (message, level) VALUES ($1, $2)`
    _, err := p.DB.Exec(query, message, level)
    if err != nil {
        return fmt.Errorf("failed to log message: %w", err)
    }
    return nil
}

// GetRecentLogs gets recent logs from database
func (p *PostgresDB) GetRecentLogs(limit int) ([]map[string]interface{}, error) {
    query := `SELECT id, message, level, created_at FROM application_logs ORDER BY created_at DESC LIMIT $1`
    
    rows, err := p.DB.Query(query, limit)
    if err != nil {
        return nil, fmt.Errorf("failed to get recent logs: %w", err)
    }
    defer rows.Close()
    
    var logs []map[string]interface{}
    for rows.Next() {
        var id int
        var message, level string
        var createdAt time.Time
        
        if err := rows.Scan(&id, &message, &level, &createdAt); err != nil {
            return nil, fmt.Errorf("failed to scan log row: %w", err)
        }
        
        log := map[string]interface{}{
            "id":         id,
            "message":    message,
            "level":      level,
            "created_at": createdAt,
        }
        logs = append(logs, log)
    }
    
    return logs, nil
}