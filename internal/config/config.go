package config

import (
	"fmt"
	"os"
)

type Config struct {
	PostgresDSN string
	HTTPPort    string
}

func LoadConfig() *Config {
	postgresHost := os.Getenv("POSTGRES_HOST")
	postgresPort := os.Getenv("POSTGRES_PORT")
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresDB := os.Getenv("POSTGRES_DB")
	
	postgresDSN := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		postgresUser, postgresPassword, postgresHost, postgresPort, postgresDB)
	
	httpPort := os.Getenv("HTTP_PORT")

	return &Config{
		PostgresDSN: postgresDSN,
		HTTPPort:    httpPort,
	}
}
