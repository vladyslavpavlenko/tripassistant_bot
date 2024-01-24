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

type testAPIRepo struct {
	App    *config.AppConfig
	Client *maps.Client
}

// NewGoogleAPIRepo creates a new instance of mapsapi.APIRepo that uses the Google Maps API.
func NewGoogleAPIRepo(client *maps.Client, app *config.AppConfig) mapsapi.APIRepo {
	return &googleAPIRepo{
		App:    app,
		Client: client,
	}
}

// NewTestingFirestoreRepo creates a new instance of a repository.DatabaseRepo intended for testing purposes.
func NewTestingFirestoreRepo(app *config.AppConfig) mapsapi.APIRepo {
	return &testAPIRepo{
		App: app,
	}
}
