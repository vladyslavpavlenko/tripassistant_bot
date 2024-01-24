package googleapirepo

import (
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/models"
)

// GetPlace returns the first place returned by the Google's Places API using Text Search.
// Note: uses message text and the name of the group for more relevant search results.
func (m *testAPIRepo) GetPlace(placeTitle string, options ...string) (models.Place, error) {
	var place models.Place

	return place, nil
}
