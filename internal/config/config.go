package config

import (
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/logger"
)

// AppConfig holds the application config
type AppConfig struct {
	AdminID int64
	Logger  *logger.MyLogger
}
