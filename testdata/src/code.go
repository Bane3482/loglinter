package src

import (
	"errors"
	"log"
	"log/slog"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Info("Приложение запущено",
		zap.String("version", "1.0"),
		zap.Int("port", 8080))
	logger.Error("Ошибка подключения", zap.Error(errors.New("db failed")))
	password := "1234"
	token := "1eafbc86f73dc"
	log.Print("aa" + password)
	slog.Info("aa" + token)
	slog.Error("smth!!")
}
