package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

type Config struct {
    DatabaseURL   string
    SessionSecret string
    Port          string
    AppEnv        string
}

func Load() *Config {
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found — reading from environment")
    }

    return &Config{
        DatabaseURL:   mustGet("DATABASE_URL"),
        SessionSecret: mustGet("SESSION_SECRET"),
        Port:          getOrDefault("PORT", "8080"),
        AppEnv:        getOrDefault("APP_ENV", "production"),
    }
}

func mustGet(key string) string {
    v := os.Getenv(key)
    if v == "" {
        log.Fatalf("Required env var %s is not set", key)
    }
    return v
}

func getOrDefault(key, fallback string) string {
    if v := os.Getenv(key); v != "" {
        return v
    }
    return fallback
}