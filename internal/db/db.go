package db

import (
    "database/sql"
    "fmt"
    "log"
    "time"

    _ "github.com/go-sql-driver/mysql"
    "packman-backend/internal/config"
)

var DB *sql.DB

// InitializeDatabase sets up the database connection pool and ensures the connection is established
func InitializeDatabase() error {
    cfg := config.GetConfig()

    // Create the MySQL data source name (DSN)
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
        cfg.DBUser,
        cfg.DBPassword,
        cfg.DBHost,
        cfg.DBPort,
        cfg.DBName,
    )

    // Open the database connection
    var err error
    DB, err = sql.Open("mysql", dsn)
    if err != nil {
        return fmt.Errorf("could not open db connection: %v", err)
    }

    // Set database connection pooling options
    DB.SetConnMaxLifetime(time.Minute * 5) // Reuse connections for 5 minutes
    DB.SetMaxOpenConns(10)                 // Allow up to 10 open connections at once
    DB.SetMaxIdleConns(5)                  // Keep up to 5 idle connections

    // Check if the connection is actually valid
    if err := DB.Ping(); err != nil {
        return fmt.Errorf("could not ping db: %v", err)
    }

    log.Println("Database connected successfully.")
    return nil
}

// CloseDatabase closes the database connection
func CloseDatabase() {
    if DB != nil {
        err := DB.Close()
        if err != nil {
            log.Fatalf("Error closing database connection: %v", err)
        } else {
            log.Println("Database connection closed.")
        }
    }
}
