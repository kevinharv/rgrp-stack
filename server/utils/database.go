package utils

import (
	"database/sql"
	"fmt"
	"log/slog"

	_ "github.com/lib/pq"
)

func CreateConnection(c *Configuration, logger *slog.Logger) (*sql.DB, error) {
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable", c.DBConfig.Username, c.DBConfig.Password, c.DBConfig.Host, c.DBConfig.Port, c.DBConfig.Database)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.Error("Failed to open connection to database")
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		logger.Error("Failed to ping database")
		return nil, err
	}

	logger.Info("Connected to database")
	return db, nil
}

func CloseConnection(db *sql.DB, logger *slog.Logger) error {
	err := db.Close()
	if err != nil {
		logger.Error("An error occurred while closing the database connection")
		return err
	}

	logger.Info("Closed connection to database")
	return nil
}
