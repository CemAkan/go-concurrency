package config

import (
	"log"
	"time"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	DBHost          string        `env:"DB_HOST" envDefault:"localhost"`
	DBUser          string        `env:"DB_USER" envDefault:"postgres"`
	DBPassword      string        `env:"DB_PASSWORD" envDefault:"password"`
	DBName          string        `env:"DB_NAME" envDefault:"shortener"`
	DBPort          int           `env:"DB_PORT" envDefault:"5432"`
	ServerPort      int           `env:"APP_PORT" envDefault:"3000"`
	Timeout         time.Duration `env:"TIMEOUT" envDefault:"5s"`
	CleanupInterval time.Duration `env:"CLICK_COUNTER_CLEANUP_INTERVAL" envDefault:"5m"`
	MaxEntries      int           `env:"CLICK_COUNTER_MAX_ENTRIES" envDefault:"10000"`
}

func LoadConfig() Config {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Failed to parse env: %v", err)
	}
	return cfg
}
