# Basic health check
curl http://localhost:8080/api/v1/health

# Detailed database info
curl http://localhost:8080/api/v1/logs

# Health Check Web Page
curl http://localhost:8080/web/health

# API Health Check
curl http://localhost:8080/api/v1/health

# log aplication startup
2024/01/20 10:30:00 Successfully connected to PostgreSQL database
2024/01/20 10:30:00 Table application_logs created/verified successfully
2024/01/20 10:30:00 Database migration completed successfully

# Log Error
2024/01/20 10:30:00 Failed to connect to database: dial tcp 127.0.0.1:5432: connect: connection refused