package api

import "github.com/vladyslavpavlenko/tripassistant_bot/internal/models"

type GoogleAPIRepo interface {
	GetPlace(placeQuery string) (models.Place, error)
}
