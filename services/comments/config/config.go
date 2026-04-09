package config

import (
	"sync"

	"github.com/vhysxl/voidspace/shared/utils/helper"
)

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
		Port:           helper.GetEnv("PORT", "8082"),
		DBConnString:   helper.GetEnv("DB_CONN", "postgres"),
		ContextTimeout: helper.GetEnvInt("CONTEXT_TIMEOUT", 10),
	}
}
