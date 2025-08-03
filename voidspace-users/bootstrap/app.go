package bootstrap

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"database/sql"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"time"
	"voidspace/users/config"
	"voidspace/users/database"

	"github.com/go-playground/validator/v10"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Application struct {
	Config                *config.Config
	DB                    *sql.DB //must using pointer from docs
	Validator             *validator.Validate
	DBContextTimeout      time.Duration
	HandlerContextTimeout time.Duration
	AccessTokenDuration   time.Duration
	RefreshTokenDuration  time.Duration
	PrivateKey            *rsa.PrivateKey
}

func App() (*Application, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found, using fallbacks")
	}

	privateKeyData, err := os.ReadFile(`C:\Users\Hp\Documents\projects\voidspace\private_key.pem`)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(privateKeyData)
	if block == nil {
		return nil, fmt.Errorf("failed to parse PEM block containing the key")
	}

	// Parse RSA private key
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse RSA private key: %v", err)
	}

	privateKey, ok := key.(*rsa.PrivateKey)
	if !ok {
		log.Fatal("key is not RSA private key")
	}

	cfg := config.GetConfig()

	var dbConfig = mysql.Config{
		User:                 cfg.DBUser,
		Passwd:               cfg.DBPassword,
		Addr:                 cfg.DBAddress,
		DBName:               cfg.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.DBContextTimeout)*time.Second)
	defer cancel()

	db, err := database.MySqlDatabase(ctx, dbConfig)
	if err != nil {
		return nil, err
	}

	return &Application{
		Config:                cfg,
		DB:                    db,
		Validator:             validator.New(),
		DBContextTimeout:      time.Duration(cfg.DBContextTimeout) * time.Second,
		HandlerContextTimeout: time.Duration(cfg.HandlerContextTimeout) * time.Second,
		AccessTokenDuration:   time.Duration(cfg.AccessTokenDuration) * time.Hour,
		RefreshTokenDuration:  time.Duration(cfg.RefreshTokenDuration) * 24 * time.Hour,
		PrivateKey:            privateKey,
	}, nil
}
