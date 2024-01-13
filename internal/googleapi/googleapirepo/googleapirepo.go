package googleapirepo

import (
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/api"
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/config"
)

type googleAPIRepo struct {
	App    *config.AppConfig
	APIKey string
}

func NewGoogleAPIRepo(APIKey string, app *config.AppConfig) api.GoogleAPIRepo {
	return &googleAPIRepo{
		App:    app,
		APIKey: APIKey,
	}
}
