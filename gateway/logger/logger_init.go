package logger

import "go.uber.org/zap"

func InitLogger() (*zap.Logger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	logger.Info("zap ready")
	return logger, nil
}
