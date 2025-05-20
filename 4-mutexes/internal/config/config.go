package config

import (
	"time"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	DBHost     string        `env:"DB_HOST" envDefault:"localhost"`
	DBUser     string        `env:"DB_USER" envDefault:"postgres"`
	DBPassword string        `env:"DB_PASSWORD" envDefault:"password"`
	DBName     string        `env:"DB_NAME" envDefault:"shortener"`
	DBPort     int           `env:"DB_PORT" envDefault:"5432"`
	ServerPort int           `env:"SERVER_PORT" envDefault:"3000"`
	Timeout    time.Duration `env:"TIMEOUT" envDefault:"5s"`
}

func Load() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
