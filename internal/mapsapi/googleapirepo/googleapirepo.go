package googleapirepo

import (
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/config"
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/mapsapi"
	"googlemaps.github.io/maps"
)

type googleAPIRepo struct {
	App    *config.AppConfig
	Client *maps.Client
}

func NewGoogleAPIRepo(client *maps.Client, app *config.AppConfig) mapsapi.APIRepo {
	return &googleAPIRepo{
		App:    app,
		Client: client,
	}
}