package config

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

// database struct
type Config struct {
	PublicHost            string
	Port                  string
	DBUser                string
	DBPassword            string
	DBAddress             string
	DBName                string
	DBContextTimeout      int
	HandlerContextTimeout int
	AccessTokenDuration   int
	RefreshTokenDuration  int
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
		PublicHost:            getEnv("PUBLIC_HOST", "http://localhost"),
		Port:                  getEnv("PORT", ":8080"),
		DBUser:                getEnv("DB_USER", "root"),
		DBPassword:            getEnv("DB_PASS", "secret"),
		DBAddress:             fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBName:                getEnv("DB_NAME", "voidspace"),
		DBContextTimeout:      getIntEnv("DB_CONTEXT_TIMEOUT", 5),
		HandlerContextTimeout: getIntEnv("H_CONTEXT_TIMEOUT", 10),
		AccessTokenDuration:   getIntEnv("ACCESS_TOKEN_DURATION", 1),
		RefreshTokenDuration:  getIntEnv("REFRESH_TOKEN_DURATION", 7),
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
