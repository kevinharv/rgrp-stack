package utils

import (
	"errors"
	"log/slog"
	"os"
	"strconv"
)

type Configuration struct {
	TLSEnabled  bool
	TLSCertPath string
	TLSKeyPath  string
	ServerPort  int
	DBConfig    *DatabaseConfig
}

type DatabaseConfig struct {
	Host     string
	Username string
	Password string
	Port     int
	Database string
}

// To-Do - Redis config type and population



func GetConfiguration() Configuration {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	logger := slog.New(handler)
	logger.Info("Reading configuration from environment")

	config := Configuration{}

	getTLSConfig(&config, logger)
	getDBConfig(&config, logger)

	return config
}



/*
Parses TLS configuration from environment variables. If variables
are missing or incorrectly set, TLS defaults to disabled.
*/
func getTLSConfig(c *Configuration, logger *slog.Logger) {
	enableTLS, configured := os.LookupEnv("TLSEnabled")
	if !configured {
		logger.Info("TLS not configured - defaulting to no TLS")
		c.TLSEnabled = false
		return
	}

	tlsEnabled, err := strconv.ParseBool(enableTLS)
	if err != nil {
		logger.Warn("Invalid value for 'TLSEnabled' - defaulting to no TLS")
		c.TLSEnabled = false
		return
	}

	if tlsEnabled {
		tlsCertPath, certConfigured := os.LookupEnv("TLSCertPath")
		tlsKeyPath, keyConfigured := os.LookupEnv("TLSKeyPath")

		if !certConfigured || !keyConfigured {
			logger.Warn("TLS Certificate and Key path not specified - defaulting to no TLS")
			c.TLSEnabled = false
		}

		if _, err := os.Stat(tlsCertPath); errors.Is(err, os.ErrNotExist) {
			logger.Warn("Certificate path does not exist - defaulting to no TLS")
			c.TLSEnabled = false
		}
		if _, err := os.Stat(tlsKeyPath); errors.Is(err, os.ErrNotExist) {
			logger.Warn("Private key path does not exist - defaulting to no TLS")
			c.TLSEnabled = false
		}

		c.TLSCertPath = tlsCertPath
		c.TLSKeyPath = tlsKeyPath

		logger.Info("Loaded TLS configuration")
	}
}



/*
Parses database configuration from environment variables. Warns
if required variables are not set or are not parsed correctly.
A valid configuration may still fail at DB connection.
*/
func getDBConfig(c *Configuration, logger *slog.Logger) {
	dbHost, hostSpecified := os.LookupEnv("DBHost")
	dbUser, userSpecified := os.LookupEnv("DBUsername")
	dbPass, passSpecified := os.LookupEnv("DBPassowrd")
	dbPort, portSpecified := os.LookupEnv("DBPort")
	dbDatabase, dbSpecified := os.LookupEnv("DBDatabase")

	if !hostSpecified || !userSpecified || !passSpecified || !portSpecified || !dbSpecified {
		logger.Error("A required database configuration parameter is missing")
		return
	}

	c.DBConfig.Host = dbHost
	c.DBConfig.Username = dbUser
	c.DBConfig.Password = dbPass
	c.DBConfig.Database = dbDatabase

	port, err := strconv.ParseInt(dbPort, 10, 32)
	if err != nil {
		logger.Error("Database port is invalid")
		return
	}

	c.DBConfig.Port = int(port)
	logger.Info("Loaded database configuration")
}
