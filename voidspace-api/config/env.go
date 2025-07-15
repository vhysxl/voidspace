package config

import (
	"fmt"
	"os"
	"sync"
)

//buat ambil config data from .env
type Config struct {
	PublicHost string
	Port			 string
	DBUser		 string
	DBPassword string
	DBAddress  string
	DBName		 string
}

var (
	envs Config
	once sync.Once //agar config di execute 1x tiap runtime	
)

func GetConfig() Config {
	once.Do(func ()  {
			envs = initConfig()  //Hanya dipanggil 1x, meski GetConfig() dipanggil berkali-kali
	})

	return envs
}

func initConfig() Config{ //init config
	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port: getEnv("PORT", "8080"),
		DBUser: getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASS", "secret"),
		DBAddress: fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"),getEnv("DB_PORT", "3306")),
		DBName: getEnv("DB_NAME", "voidspace"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
