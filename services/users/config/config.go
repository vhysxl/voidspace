package config

import (
	"os"
	"strconv"
	"sync"
)

// database struct
type Config struct {
	Port                 string
	ContextTimeout       int
	AccessTokenDuration  int
	RefreshTokenDuration int
	DBConnectionString   string
}

var (
	envs Config
	once sync.Once //run once per runtime
)

func GetConfig() *Config {
	once.Do(func() {
		envs = initConfig()
	})

	return &envs //return pointer so envs are One Source of Truth (singleton)
}

func initConfig() Config {
	return Config{
		Port:                 getEnv("PORT", ":8080"),
		ContextTimeout:       getIntEnv("CONTEXT_TIMEOUT", 10),
		AccessTokenDuration:  getIntEnv("ACCESS_TOKEN_DURATION", 30),
		RefreshTokenDuration: getIntEnv("REFRESH_TOKEN_DURATION", 7),
		DBConnectionString:   getEnv("DB_CONNECTION", "mysqlconn"),
	}
}
func getEnv(key, fallback string) string { //lookup env
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback //use fallback not from .env
}

func getIntEnv(key string, fallback int) int { //to parse env to int
	if value, ok := os.LookupEnv(key); ok {
		if value, err := strconv.Atoi(value); err == nil {
			return value
		}
	}

	return fallback
}
