package config

import (
	"log"
	"os"
	"strconv"
	"sync"
)

// database struct
type Config struct {
	PublicHost           string
	Port                 string
	DBConnectionString   string
	ContextTimeout       int
	AccessTokenDuration  int
	RefreshTokenDuration int
	SecretPath           string
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
		DBConnectionString:   getEnv("DB_CONN", "postgres"),
		ContextTimeout:       getIntEnv("CONTEXT_TIMEOUT", 10),
		AccessTokenDuration:  getIntEnv("ACCESS_TOKEN_DURATION", 30),
		RefreshTokenDuration: getIntEnv("REFRESH_TOKEN_DURATION", 7),
		SecretPath:           getEnv("SECRET_PATH", "SECRETS"),
	}
}
func getEnv(key, fallback string) string { //lookup env
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	log.Println("Using fallback for value", key)

	return fallback //use fallback not from .env
}

func getIntEnv(key string, fallback int) int { //to parse env to int
	if value, ok := os.LookupEnv(key); ok {
		if value, err := strconv.Atoi(value); err == nil {
			return value
		}
	}

	log.Println("Using fallback for value", key)

	return fallback
}
