package config

import (
	"os"
	"strconv"
	"sync"
)

// database struct
type Config struct {
	PublicHost           string
	Port                 string
	DBUser               string
	DBPassword           string
	DBAddress            string
	DBName               string
	ContextTimeout       int
	AccessTokenDuration  int
	RefreshTokenDuration int
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
		PublicHost:           getEnv("PUBLIC_HOST", "http://localhost"),
		Port:                 getEnv("PORT", ":8080"),
		DBUser:               getEnv("PROD_DB_USER", "root"),
		DBPassword:           getEnv("PROD_DB_PASS", "secret"),
		DBAddress:            getEnv("PROD_DB_ADDRESS", "localhost:3306"),
		DBName:               getEnv("PROD_DB_NAME", "voidspace"),
		ContextTimeout:       getIntEnv("CONTEXT_TIMEOUT", 10),
		AccessTokenDuration:  getIntEnv("ACCESS_TOKEN_DURATION", 30),
		RefreshTokenDuration: getIntEnv("REFRESH_TOKEN_DURATION", 7),
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
