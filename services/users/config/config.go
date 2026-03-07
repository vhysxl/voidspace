package config

import (
	"sync"

	"github.com/vhysxl/voidspace/shared/utils/helper"
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
		Port:                 helper.GetEnv("PORT", ":8080"),
		DBConnectionString:   helper.GetEnv("DB_CONN", "postgres"),
		ContextTimeout:       helper.GetEnvInt("CONTEXT_TIMEOUT", 10),
		AccessTokenDuration:  helper.GetEnvInt("ACCESS_TOKEN_DURATION", 30),
		RefreshTokenDuration: helper.GetEnvInt("REFRESH_TOKEN_DURATION", 7),
		SecretPath:           helper.GetEnv("SECRET_PATH", "SECRETS"),
	}
}
