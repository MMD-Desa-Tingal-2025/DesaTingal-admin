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

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/fachru/backend/config"
	"github.com/fachru/backend/database"
	"github.com/fachru/backend/handlers"
	"github.com/fachru/backend/middleware"

	_ "github.com/lib/pq"
    _ "github.com/golang-migrate/migrate/v4/source/file"
)

// Tambahkan fungsi ini untuk cek koneksi DB
func CheckDBConnection(db *sql.DB) error {
	var now time.Time
	err := db.QueryRow("SELECT NOW()").Scan(&now)
	if err != nil {
		return err
	}
	log.Printf("PostgreSQL connected. Server time: %v", now)
	return nil
}

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	// Initialize database
	db, err := database.NewPostgresDB(cfg.Database.URL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Jalankan test koneksi PostgreSQL
	if err := CheckDBConnection(db.DB); err != nil {
		log.Fatalf("Failed to test DB connection: %v", err)
	}

	// Run database migration
	if err := db.RunSQLMigration(); err != nil {
		log.Printf("Warning: Migration failed: %v", err)
	}

	// Log successful startup
	if err := db.LogMessage("Application started successfully", "info"); err != nil {
		log.Printf("Failed to log startup: %v", err)
	}

	// Setup Gin router
	if cfg.App.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	// Add CORS middleware
	router.Use(middleware.CORSMiddleware())

	// Initialize handlers
	apiHandler := handlers.NewAPIHandler(db)

	// Serve static files for production
	if cfg.App.Environment == "production" {
		router.Static("/static", "../frontend/dist/static")
		router.StaticFile("/", "../frontend/dist/index.html")
	}

	// API routes
	api := router.Group("/api/v1")
	{
		api.GET("/health", apiHandler.HealthCheck)
		api.GET("/logs", apiHandler.GetLogs)
		api.POST("/logs", apiHandler.CreateLog)
		api.GET("/dashboard/data", apiHandler.GetDashboardData)
		api.GET("/map/data", apiHandler.GetMapData)
	}

	// HTTP server config
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: router,
	}

	// Run server
	go func() {
		log.Printf("Starting server on port %d...\n", cfg.Server.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Wait for shutdown signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Log shutdown
	if err := db.LogMessage("Application shutting down", "info"); err != nil {
		log.Printf("Failed to log shutdown: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")


}
