package googleapi

import "github.com/vladyslavpavlenko/tripassistant_bot/internal/models"

type APIRepo interface {
	GetPlace(placeQuery string) (models.Place, error)
}
