package config

import (
	"crypto/rsa"
	"fmt"
	"os"
	"strconv"
	"sync"
	"voidspaceGateway/utils"
)

type Config struct {
	Port            string
	PublicKey       *rsa.PublicKey
	ApiSecret       string
	ContextTimeout  int
	UserServiceAddr string
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
	publicKey, err := utils.LoadPublicKey(getEnv("PUBLIC_KEY_PATH", "./certs/public_key.pem"))
	if err != nil {
		fmt.Println("FAILED TO LOAD PUBLIC KEY FROM ENV AND FALLBACK!")
		panic("error load the public key")
	}

	return Config{
		Port:            getEnv("PORT", ":5000"),
		PublicKey:       publicKey,
		ApiSecret:       getEnv("I_SECRET", "SUPER SECRET LMAO"),
		ContextTimeout:  getIntEnv("CONTEXT_TIMEOUT", 10),
		UserServiceAddr: getEnv("USER_SERVICE_URL", "localhost:8080"),
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
