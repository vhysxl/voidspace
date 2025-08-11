package config

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

// Config struct
type Config struct {
	PublicHost     string
	Port           string
	DBUser         string
	DBPassword     string
	DBHost         string
	DBPort         string
	DBName         string
	DBSSLMode      string
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
		PublicHost:     getEnv("PUBLIC_HOST", "http://localhost"),
		Port:           getEnv("PORT", "8080"),
		DBUser:         getEnv("DB_USER", "postgres"),
		DBPassword:     getEnv("DB_PASSWORD", "secret"),
		DBHost:         getEnv("DB_HOST", "localhost"),
		DBPort:         getEnv("DB_PORT", "5432"),
		DBName:         getEnv("DB_NAME", "voidspace"),
		DBSSLMode:      getEnv("DB_SSLMODE", "disable"),
		ContextTimeout: getIntEnv("CONTEXT_TIMEOUT", 10),
	}
}

// GetDBConnectionString returns PostgreSQL connection string
func (c *Config) GetDBConnectionString() string {
	return fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		c.DBUser,
		c.DBPassword,
		c.DBHost,
		c.DBPort,
		c.DBName,
		c.DBSSLMode,
	)
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
