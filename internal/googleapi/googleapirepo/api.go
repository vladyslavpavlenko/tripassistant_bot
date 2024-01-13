package googleapirepo

import (
	"context"
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/models"
	"googlemaps.github.io/maps"
	"time"
)

const requestTimeout = 3 * time.Second

// GetPlace returns the first place returned by the Google's Places API using Text Search
func (m *googleAPIRepo) GetPlace(placeQuery string) (models.Place, error) {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	var place models.Place

	r := &maps.TextSearchRequest{
		Query: placeQuery,
	}

	response, err := m.Client.TextSearch(ctx, r)
	if err != nil {
		return place, err
	}

	if len(response.Results) > 0 {
		firstPlace := response.Results[0]

		place.PlaceTitle = placeQuery
		place.PlaceLatitude = firstPlace.Geometry.Location.Lat
		place.PlaceLongitude = firstPlace.Geometry.Location.Lng
		place.PlaceAddress = firstPlace.FormattedAddress
	}

	return place, nil
}
