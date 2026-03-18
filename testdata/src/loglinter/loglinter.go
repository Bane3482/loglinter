package loglinter

import (
	"log/slog"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewDevelopment()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			panic(err)
		}
	}(logger)
	password := "secret"
	apiKey := "secretKey"
	token := "secretToken"

	// Zap messages
	logger.Info("Starting server on port 8080")
	logger.Error("Failed to connect to database")
	logger.Info("запуск сервера")
	logger.Error("ошибка подключения к базе данных")
	logger.Info("server started!🚀")
	logger.Error("connection failed!!!")
	logger.Warn("warning: something went wrong...")
	logger.Info("user password " + password)
	logger.Debug("api_key=" + apiKey)
	logger.Info("token " + token)

	logger.Info("starting server on port 8080")
	logger.Error("failed to connect to database")
	logger.Info("starting server")
	logger.Error("failed to connect to database")
	logger.Info("server started")
	logger.Error("connection failed")
	logger.Warn("something went wrong")
	logger.Info("server started")
	logger.Error("connection failed")
	logger.Warn("something went wrong")
	logger.Info("user authenticated successfully")
	logger.Debug("api request completed")
	logger.Info("token validated")

	// Slog messages
	slog.Info("Starting server on port 8080")
	slog.Error("Failed to connect to database")
	slog.Info("запуск сервера")
	slog.Error("ошибка подключения к базе данных")
	slog.Info("server started!🚀")
	slog.Error("connection failed!!!")
	slog.Warn("warning: something went wrong...")
	slog.Info("user password: " + password)
	slog.Debug("api_key" + apiKey)
	slog.Info("token " + token)

	slog.Info("starting server on port 8080")
	slog.Error("failed to connect to database")
	slog.Info("starting server")
	slog.Error("failed to connect to database")
	slog.Info("server started")
	slog.Error("connection failed")
	slog.Warn("something went wrong")
	slog.Info("server started")
	slog.Error("connection failed")
	slog.Warn("something went wrong")
	slog.Info("user authenticated successfully")
	slog.Debug("api request completed")
	slog.Info("token' validated")
}
