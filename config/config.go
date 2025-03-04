package config

type Config struct {
	Driver string
	DSN    string
}

func NewConfig() *Config {
	driver := "postgres"
	dsn := "host=localhost user=postgres password=1234 dbname=testdb port=5432 sslmode=disable"

	return &Config{
		Driver: driver,
		DSN:    dsn,
	}
}
