package config

import "os"

type Config struct {
	DSN  string
	Port string
	Mode string
}

func PrepareConfig() *Config {
	return &Config{
		DSN:  os.Getenv("DB_SOURCE"),
		Port: os.Getenv("APP_PORT"),
		Mode: os.Getenv("APP_MODE"),
	}
}
