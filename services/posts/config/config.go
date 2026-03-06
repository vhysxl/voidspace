package config

import (
	"os"
	"strconv"
	"sync"
)

// Config struct
type Config struct {
	Port           string
	DBConnString   string
	ContextTimeout int
}

var (
	envs Config
	once sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		envs = initConfig()
	})

	return &envs
}

func initConfig() Config {
	return Config{
		Port:           getEnv("PORT", "8080"),
		DBConnString:   getEnv("DB_CONN", "postgres"),
		ContextTimeout: getIntEnv("CONTEXT_TIMEOUT", 10),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getIntEnv(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return fallback
}
