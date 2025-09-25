package config

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

type Config struct {
	PublicHost             string
	Port                   string
	DBUser                 string
	DBPassword             string
	DBAddress              string
	DBName                 string
	ContextTimeout         int
	InstanceConnectionName string
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
		PublicHost:             getEnv("PUBLIC_HOST", "http://localhost"),
		Port:                   getEnv("PORT", ":8082"),
		DBUser:                 getEnv("PROD_DB_USER", "root"),
		DBPassword:             getEnv("PROD_DB_PASS", "secret"),
		DBAddress:              fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3307")),
		DBName:                 getEnv("PROD_DB_NAME", "voidspace"),
		ContextTimeout:         getIntEnv("CONTEXT_TIMEOUT", 10),
		InstanceConnectionName: getEnv("PROD_INSTANCE_CONNECTION_NAME", "project:region:instance"),
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
		if value, err := strconv.Atoi(value); err == nil {
			return value
		}
	}

	return fallback
}
