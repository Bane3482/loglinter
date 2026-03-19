package loglinter

import (
	"log/slog"

	"go.uber.org/zap"
)

func main() {
	logger := &zap.Logger{}

	password := "secret"
	apiKey := "secretKey"
	token := "secretToken"

	// Zap messages
	logger.Info("Starting server on port 8080")      // want `message starts not from small letter`
	logger.Error("Failed to connect to database")    // want `Incorrect message`
	logger.Info("запуск сервера")                    // want `Message contains not english text`
	logger.Error("ошибка подключения к базе данных") // want `Message contains not english text`
	logger.Info("server started!🚀")                  // want `Message contains special symbols`
	logger.Error("connection failed!!!")             // want `Message contains special symbols`
	logger.Warn("warning: something went wrong...")  // want `Message contains special symbols`
	logger.Info("user password " + password)         // want `Message contains potencial sensitive data`
	logger.Debug("api_key=" + apiKey)                // want `Message contains special symbols`
	logger.Info("token " + token)                    // want `Message contains potencial sensitive data`

	logger.Info("starting server on port 8080")    // ok
	logger.Error("failed to connect to database")  // ok
	logger.Info("starting server")                 // ok
	logger.Error("failed to connect to database")  // ok
	logger.Info("server started")                  // ok
	logger.Error("connection failed")              // ok
	logger.Warn("something went wrong")            // ok
	logger.Info("server started")                  // ok
	logger.Error("connection failed")              // ok
	logger.Warn("something went wrong")            // ok
	logger.Info("user authenticated successfully") // ok
	logger.Debug("api request completed")          // ok
	logger.Info("token validated")                 // ok

	// Slog messages
	slog.Info("Starting server on port 8080")      // want `Message starts not from small letter`
	slog.Error("Failed to connect to database")    // want `Message starts not from small letter`
	slog.Info("запуск сервера")                    // want `Message contains not english text`
	slog.Error("ошибка подключения к базе данных") // want `Message contains not english text`
	slog.Info("server started!🚀")                  // want `Message contains special symbols`
	slog.Error("connection failed!!!")             // want `Message contains special symbols`
	slog.Warn("warning: something went wrong...")  // want `Message contains special symbols`
	slog.Info("user password: " + password)        // want "Message contains special symbols"
	slog.Debug("api_key" + apiKey)                 // want `Message contains special symbols`
	slog.Info("token " + token)                    // want "Message contains potencial sensitive data"

	slog.Info("starting server on port 8080")    // ok
	slog.Error("failed to connect to database")  // ok
	slog.Info("starting server")                 // ok
	slog.Error("failed to connect to database")  // ok
	slog.Info("server started")                  // ok
	slog.Error("connection failed")              // ok
	slog.Warn("something went wrong")            // ok
	slog.Info("server started")                  // ok
	slog.Error("connection failed")              // ok
	slog.Warn("something went wrong")            // ok
	slog.Info("user authenticated successfully") // ok
	slog.Debug("api request completed")          // ok
	slog.Info("token' validated")                // want "Message contains special symbols"
}
