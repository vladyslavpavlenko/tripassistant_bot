package mapsapi

import "github.com/vladyslavpavlenko/tripassistant_bot/internal/models"

type APIRepo interface {
	GetPlace(placeTitle string, options ...string) (models.Place, error)
}
