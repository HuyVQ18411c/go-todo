package config

import "os"

type Config struct {
	// App address config
	Address string
	// Database config
	DatabaseHost     string
	DatabasePort     string
	DatabaseUser     string
	DatabasePassword string
}

func (config *Config) GetDatabaseURL(useSQLite bool) string {
	if useSQLite {
		return "test.db"
	}
	return ""
}

func NewConfig() *Config {
	config := Config{
		Address:          os.Getenv("APP_ADDRESS"),
		DatabaseHost:     os.Getenv("DATABASE_HOST"),
		DatabasePort:     os.Getenv("DATABASE_PORT"),
		DatabaseUser:     os.Getenv("DATABASE_USER"),
		DatabasePassword: os.Getenv("DATABASE_PASSWORD"),
	}
	return &config
}
