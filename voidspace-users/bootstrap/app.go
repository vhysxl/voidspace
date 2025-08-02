package bootstrap

import (
	"context"
	"database/sql"
	"log"
	"time"
	"voidspace/users/config"
	"voidspace/users/database"

	"github.com/go-playground/validator/v10"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Application struct {
	Config         *config.Config
	DB             *sql.DB //must using pointer from docs
	Validator      *validator.Validate
	ContextTimeout time.Duration
}

func App() (*Application, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found, using fallbacks")
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.ContextTimeout)*time.Second)
	defer cancel()

	db, err := database.MySqlDatabase(ctx, dbConfig)
	if err != nil {
		return nil, err
	}

	return &Application{
		Config:         cfg,
		DB:             db,
		Validator:      validator.New(),
		ContextTimeout: time.Duration(cfg.ContextTimeout) * time.Second,
	}, nil
}
