package config

import (
	"crypto/rsa"
	"fmt"
	"sync"
	"voidspaceGateway/utils"

	"github.com/vhysxl/voidspace/shared/utils/helper"
)

type Config struct {
	Port               string
	PublicKey          *rsa.PublicKey
	ApiSecret          string
	ContextTimeout     int
	UserServiceAddr    string
	PostServiceAddr    string
	CommentServiceAddr string
	BucketName         string
	TemporalPort       string
	Environment        string
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
	publicKey, err := utils.LoadPublicKey(helper.GetEnv("PUBLIC_KEY_PATH", "/etc/secrets/public-key"))
	if err != nil {
		fmt.Println("FAILED TO LOAD PUBLIC KEY FROM ENV AND FALLBACK!")
		panic("error load the public key")
	}

	return Config{
		Port:               helper.GetEnv("PORT", "8080"),
		PublicKey:          publicKey,
		ApiSecret:          helper.GetEnv("API_SECRET", "SUPER SECRET LMAO"),
		ContextTimeout:     helper.GetEnvInt("CONTEXT_TIMEOUT", 30),
		UserServiceAddr:    helper.GetEnv("USER_SERVICE_URL", "localhost:8080"),
		PostServiceAddr:    helper.GetEnv("POST_SERVICE_URL", "localhost:5000"),
		CommentServiceAddr: helper.GetEnv("COMMENT_SERVICE_URL", "localhost:8082"),
		BucketName:         helper.GetEnv("BUCKET_NAME", "assets_voidspace"),
		TemporalPort:       helper.GetEnv("TEMPORAL_PORT", "localhost:7233"),
		Environment:        helper.GetEnv("ENV", "PROD"),
	}
}
