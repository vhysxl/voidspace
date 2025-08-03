package bootstrap

import (
	"context"
	"crypto/rsa"
	"database/sql"
	"log"
	"time"
	"voidspace/users/config"
	"voidspace/users/database"
	"voidspace/users/logger"

	"github.com/go-playground/validator/v10"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
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
	Logger                *zap.Logger
}

func App() (*Application, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found, using fallbacks", err)
	}

	logger, err := logger.InitLogger()
	if err != nil {
		log.Println("logger failed to load", err)
	}
	defer logger.Sync()

	privateKey, err := config.LoadPrivateKey("./secret/private_key.pem")
	if err != nil {
		logger.Error("Failed to load private key", zap.Error(err))
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
		logger.Error("Failed to connect to database", zap.Error(err))
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
		Logger:                logger,
	}, nil
}
