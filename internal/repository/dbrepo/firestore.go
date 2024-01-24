package dbrepo

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/models"
	"strconv"
	"time"
)

const (
	requestTimeout = 3 * time.Second
)

// AddUser adds a user to the users collection.
func (m *firestoreDBRepo) AddUser(user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	docName := strconv.FormatInt(user.UserID, 10)

	docRef := m.Client.Collection("users").Doc(docName)
	_, err := docRef.Set(ctx, map[string]any{
		"user_id":   user.UserID,
		"user_name": user.UserName,
	})
	if err != nil {
		return err
	}

	return nil
}

// GetAllUserIDs returns all the user IDs from users collection.
func (m *firestoreDBRepo) GetAllUserIDs() ([]int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	var userIDs []int64

	docRef := m.Client.Collection("users")

	documents, err := docRef.Documents(ctx).GetAll()
	if err != nil {
		return userIDs, err
	}

	for _, doc := range documents {
		var data map[string]interface{}
		if err := doc.DataTo(&data); err != nil {
			fmt.Println("Error converting document data:", err)
			continue
		}

		if id, exists := data["user_id"]; exists {
			idInt64 := id.(int64)
			userIDs = append(userIDs, idInt64)
		}
	}

	return userIDs, nil
}

// DeleteUserByID deletes a user from the users collection based on its ID.
func (m *firestoreDBRepo) DeleteUserByID(userID int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	docName := strconv.FormatInt(userID, 10)

	docRef := m.Client.Collection("users").Doc(docName)
	_, err := docRef.Delete(ctx)
	if err != nil {
		return err
	}

	return nil
}

// CheckIfUserIsRegisteredByID checks whether a user is already registered in the users collection by their ID.
func (m *firestoreDBRepo) CheckIfUserIsRegisteredByID(userID int64) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	usersCollection := m.Client.Collection("users")

	docName := strconv.FormatInt(userID, 10)

	docRef := usersCollection.Doc(docName)
	docSnapshot, err := docRef.Get(ctx)
	if err != nil {
		return false, err
	}

	return docSnapshot.Exists(), nil
}

// AddTrip adds a trip to the trips collection.
// Note: trip is a group chat.
func (m *firestoreDBRepo) AddTrip(trip models.Trip) error {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	docName := strconv.FormatInt(trip.TripID, 10)

	docRef := m.Client.Collection("trips").Doc(docName)
	_, err := docRef.Set(ctx, map[string]any{
		"trip_id":     trip.TripID,
		"trip_places": trip.TripPlaces,
	})
	if err != nil {
		return err
	}

	return nil
}

// DeleteTripByID deletes a trip from the trips collection based on its ID.
// Note: trip is a group chat.
func (m *firestoreDBRepo) DeleteTripByID(tripID int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	docName := strconv.FormatInt(tripID, 10)

	docRef := m.Client.Collection("trips").Doc(docName)
	_, err := docRef.Delete(ctx)
	if err != nil {
		return err
	}

	return nil
}

// AddPlaceToListByTripID adds the place to the list of the trip identified by its ID.
// Note: trip is a group chat.
func (m *firestoreDBRepo) AddPlaceToListByTripID(place models.Place, tripID int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*requestTimeout)
	defer cancel()

	oldTripPlaces, err := m.GetTripPlacesListByID(tripID)
	if err != nil {
		return err
	}

	tripPlaces := append(oldTripPlaces, place)

	docName := strconv.FormatInt(tripID, 10)

	_, err = m.Client.Collection("trips").Doc(docName).Update(ctx, []firestore.Update{
		{Path: "trip_places", Value: tripPlaces},
	})
	if err != nil {
		return err
	}

	return nil
}

// GetTripPlacesListByID returns all the places from the trip by its ID.
// Note: trip is a group chat.
func (m *firestoreDBRepo) GetTripPlacesListByID(tripID int64) ([]models.Place, error) {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	var tripPlaces []models.Place

	docName := strconv.FormatInt(tripID, 10)

	doc, err := m.Client.Collection("trips").Doc(docName).Get(ctx)
	if err != nil {
		return tripPlaces, err
	}

	data := doc.Data()

	placesData := data["trip_places"].([]any)

	for _, placeData := range placesData {
		placeMap, ok := placeData.(map[string]interface{})
		if !ok {
			return nil, err
		}

		place := models.Place{
			PlaceID:        placeMap["place_id"].(string),
			PlaceTitle:     placeMap["place_title"].(string),
			PlaceLatitude:  placeMap["place_latitude"].(float64),
			PlaceLongitude: placeMap["place_longitude"].(float64),
			PlaceAddress:   placeMap["place_address"].(string),
		}

		tripPlaces = append(tripPlaces, place)
	}

	return tripPlaces, nil
}

// GetTripByID returns the trip by its ID.
// Note: trip is a group chat.
func (m *firestoreDBRepo) GetTripByID(tripID int64) (models.Trip, error) {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	var trip models.Trip
	var tripPlaces []models.Place

	docName := strconv.FormatInt(tripID, 10)

	doc, err := m.Client.Collection("trips").Doc(docName).Get(ctx)
	if err != nil {
		return trip, err
	}

	data := doc.Data()

	placesData := data["trip_places"].([]any)

	for _, placeData := range placesData {
		placeMap, ok := placeData.(map[string]interface{})
		if !ok {
			return trip, err
		}

		place := models.Place{
			PlaceID:        placeMap["place_id"].(string),
			PlaceTitle:     placeMap["place_title"].(string),
			PlaceLatitude:  placeMap["place_latitude"].(float64),
			PlaceLongitude: placeMap["place_longitude"].(float64),
			PlaceAddress:   placeMap["place_address"].(string),
		}

		tripPlaces = append(tripPlaces, place)
	}

	trip.TripID = tripID
	trip.TripPlaces = tripPlaces

	return trip, nil
}

// DeleteTripPlacesListByID deletes trip places list by its ID.
// Note: trip is a group chat.
func (m *firestoreDBRepo) DeleteTripPlacesListByID(tripID int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	docName := strconv.FormatInt(tripID, 10)

	_, err := m.Client.Collection("trips").Doc(docName).Update(ctx, []firestore.Update{
		{
			Path:  "trip_places",
			Value: []models.Place{},
		},
	})
	if err != nil {
		return err
	}

	return nil
}

// DeleteTripPlaceByTitle deletes a place from trip_places array by its title.
// Note: trip is a group chat.
func (m *firestoreDBRepo) DeleteTripPlaceByTitle(placeTitle string, tripID int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	docName := strconv.FormatInt(tripID, 10)

	docRef := m.Client.Collection("trips").Doc(docName)
	doc, err := docRef.Get(ctx)
	if err != nil {
		return err
	}

	var tripData map[string]any
	err = doc.DataTo(&tripData)
	if err != nil {
		return err
	}

	tripPlaces := tripData["trip_places"].([]any)

	updatedTripPlaces := make([]any, 0, len(tripPlaces))
	for _, place := range tripPlaces {
		if placeMap, isMap := place.(map[string]any); isMap {
			if name, exists := placeMap["place_title"].(string); exists && name == placeTitle {
				continue
			}
			updatedTripPlaces = append(updatedTripPlaces, placeMap)
		}
	}

	_, err = docRef.Update(ctx, []firestore.Update{
		{
			Path:  "trip_places",
			Value: updatedTripPlaces,
		},
	})
	if err != nil {
		return err
	}

	return nil
}
