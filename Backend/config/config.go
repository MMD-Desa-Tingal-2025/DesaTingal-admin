// backend/config/config.go
package config

import (
    "fmt"
    "os"
    "strconv"
)

type Config struct {
    App      AppConfig
    Database DatabaseConfig
    Server   ServerConfig
}

type AppConfig struct {
    Environment string
}

type DatabaseConfig struct {
    URL string
}

type ServerConfig struct {
    Port int
}

// Load loads configuration from environment variables
func Load() (*Config, error) {
    port, err := strconv.Atoi(getEnv("SERVER_PORT", "8080"))
    if err != nil {
        return nil, fmt.Errorf("invalid SERVER_PORT: %w", err)
    }

    // Build database URL from environment variables
    dbURL := buildDatabaseURL()

    config := &Config{
        App: AppConfig{
            Environment: getEnv("APP_ENV", "development"),
        },
        Database: DatabaseConfig{
            URL: dbURL,
        },
        Server: ServerConfig{
            Port: port,
        },
    }

    return config, nil
}

// buildDatabaseURL constructs PostgreSQL connection URL from environment variables
func buildDatabaseURL() string {
    host := getEnv("DB_HOST", "localhost")
    port := getEnv("DB_PORT", "5432")
    user := getEnv("DB_USER", "postgres")
    password := getEnv("DB_PASSWORD", "")
    dbname := getEnv("DB_NAME", "myapp")
    sslmode := getEnv("DB_SSLMODE", "disable")

    return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
        user, password, host, port, dbname, sslmode)
}

// getEnv gets environment variable with fallback
func getEnv(key, fallback string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return fallback
}