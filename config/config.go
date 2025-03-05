package config

import "os"

type Config struct {
	Driver string
	DSN    string
}

func NewConfig() *Config {
	driver := "postgres"
	dsn := os.Getenv("DSN")

	return &Config{
		Driver: driver,
		DSN:    dsn,
	}
}
