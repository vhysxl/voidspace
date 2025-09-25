package config

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/jackc/pgx/v5"
)

// Config struct
type Config struct {
	Port                   string
	DBUser                 string
	DBPassword             string
	DBName                 string
	DBSSLMode              string
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
		Port:                   getEnv("PORT", "8081"),
		DBUser:                 getEnv("PROD_DB_USER", "postgres"),
		DBPassword:             getEnv("PROD_DB_PASS", "secret"),
		DBName:                 getEnv("PROD_DB_NAME", "voidspace"),
		ContextTimeout:         getIntEnv("CONTEXT_TIMEOUT", 10),
		InstanceConnectionName: getEnv("PROD_INSTANCE_CONNECTION_NAME", "project:region:instance"),
	}
}

// GetDBConnectionString returns PostgreSQL connection string
func (c *Config) GetDBConnectionString() (*pgx.ConnConfig, error) {
	dsn := fmt.Sprintf("user=%s password=%s database=%s", c.DBUser, c.DBPassword, c.DBName)
	config, err := pgx.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	return config, nil
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
