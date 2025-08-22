package bootstrap

import (
	"log"
	"time"
	"voidspaceGateway/config"
	logger "voidspaceGateway/loggger"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type Application struct {
	Config         *config.Config
	ContextTimeout time.Duration
	Validator      *validator.Validate
	Logger         *zap.Logger
}

func App() (*Application, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found, using fallbacks", err)
	}

	config := config.GetConfig()

	validator := validator.New()

	logger, err := logger.InitLogger()
	if err != nil {
		log.Println("Logger failed to load", err)
		return nil, err
	}

	logger.Info("Gateway Ready")
	return &Application{
		Config:         config,
		ContextTimeout: time.Duration(config.ContextTimeout),
		Validator:      validator,
		Logger:         logger,
	}, nil

}
