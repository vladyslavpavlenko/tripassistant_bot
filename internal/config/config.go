package config

import (
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/logger"
)

// AppConfig holds the application config.
type AppConfig struct {
	AdminIDs    []int64
	Logger      *logger.MyLogger
	BotUsername string
}
