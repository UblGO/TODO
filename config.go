package main

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Port     string
	DB       PostgresConfig
	CreateDB string
}

type PostgresConfig struct {
	Username string
	Password string
	URL      string
	Port     string
	DbName   string
}

// Возвращает строку для подключения к Postgres
func (p *PostgresConfig) Pgconn() string {
	return fmt.Sprintf("user=%s password=%s host=%s sslmode=disable", p.Username, p.Password, p.URL)
}

// Возвращает строку для подключения к БД
func (p *PostgresConfig) ConnInfo() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", p.Username, p.Password, p.URL, p.Port, p.DbName)
}

func LoadConfig() *Config {
	cfg := &Config{
		Port: os.Getenv("PORT"),
		DB: PostgresConfig{
			Username: os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			URL:      os.Getenv("POSTGRES_URL"),
			Port:     os.Getenv("POSTGRES_PORT"),
			DbName:   os.Getenv("POSTGRES_DB"),
		},
		CreateDB: os.Getenv("CREATE_DB"),
	}

	return cfg
}
