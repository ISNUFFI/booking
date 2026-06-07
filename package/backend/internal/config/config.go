package config

import (
	"os"
	"log"
)

type Config struct {
	AppAddress string
	DatabaseURL string
	JWTSecret string
}

func Load() *Config {
	return &Config {
		AppAddress: getEnv("APP_ADDRESS"),
		DatabaseURL: getEnv("DATABASE_URL"),
		JWTSecret: getEnv("JWT_SECRET"),
	}
}

func getEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("missing env: ", key)
	}

	return v
}
