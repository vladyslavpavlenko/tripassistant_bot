package googleapirepo

import (
	"context"
	"errors"
	"fmt"
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/models"
	"googlemaps.github.io/maps"
	"time"
)

const requestTimeout = 3 * time.Second

// GetPlace returns the first place returned by the Google's Places API using Text Search.
// Note: uses message text and the name of the group for more relevant search results.
func (m *googleAPIRepo) GetPlace(placeTitle string, options ...string) (models.Place, error) {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	var place models.Place
	query := placeTitle

	if len(options) != 0 {
		for _, i := range options {
			query += fmt.Sprintf(" %s", i)
		}
	}

	r := &maps.TextSearchRequest{
		Query: query,
	}

	response, err := m.Client.TextSearch(ctx, r)
	if err != nil {
		return place, err
	}

	if len(response.Results) > 0 {
		firstPlace := response.Results[0]

		place.PlaceID = firstPlace.PlaceID
		place.PlaceTitle = placeTitle
		place.PlaceLatitude = firstPlace.Geometry.Location.Lat
		place.PlaceLongitude = firstPlace.Geometry.Location.Lng
		place.PlaceAddress = firstPlace.FormattedAddress
	} else {
		return place, errors.New("couldn't find the place")
	}

	return place, nil
}
