package dbrepo

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/models"
	"strconv"
	"time"
)

const (
	requestTimeout      = 3 * time.Second
	requestDebugTimeout = 600 * time.Second
)

// AddUser adds a user to the users collection
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

// DeleteUserByID deletes a user from the users collection based on its ID
func (m *firestoreDBRepo) DeleteUserByID(id int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	docName := strconv.FormatInt(id, 10)

	docRef := m.Client.Collection("users").Doc(docName)
	_, err := docRef.Delete(ctx)
	if err != nil {
		return err
	}

	return nil
}

// CheckIfUserIsRegisteredByID checks whether a user is already registered in the users collection by their ID
func (m *firestoreDBRepo) CheckIfUserIsRegisteredByID(id int64) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	usersCollection := m.Client.Collection("users")

	docName := strconv.FormatInt(id, 10)

	docRef := usersCollection.Doc(docName)
	docSnapshot, err := docRef.Get(ctx)
	if err != nil {
		return false, err
	}

	return docSnapshot.Exists(), nil
}

// AddTrip adds a trip to the trips collection
// Note: trip is a group chat
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

// DeleteTripByID deletes a trip from the trips collection based on its ID
// Note: trip is a group chat
func (m *firestoreDBRepo) DeleteTripByID(id int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	docName := strconv.FormatInt(id, 10)

	docRef := m.Client.Collection("trips").Doc(docName)
	_, err := docRef.Delete(ctx)
	if err != nil {
		return err
	}

	return nil
}

// AddPlaceToListByTripID adds the place to the list of the trip identified by its ID
// Note: trip is a group chat
func (m *firestoreDBRepo) AddPlaceToListByTripID(place models.Place, id int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*requestTimeout)
	defer cancel()

	oldTripPlaces, err := m.GetTripPlacesListByID(id)
	if err != nil {
		return err
	}

	tripPlaces := append(oldTripPlaces, place)

	docName := strconv.FormatInt(id, 10)

	_, err = m.Client.Collection("trips").Doc(docName).Update(ctx, []firestore.Update{
		{Path: "trip_places", Value: tripPlaces},
	})
	if err != nil {
		return err
	}

	return nil
}

// GetTripPlacesListByID returns all the places from the trip by its ID
// Note: trip is a group chat
func (m *firestoreDBRepo) GetTripPlacesListByID(id int64) ([]models.Place, error) {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	var tripPlaces []models.Place

	docName := strconv.FormatInt(id, 10)

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
			PlaceTitle:     placeMap["place_title"].(string),
			PlaceLatitude:  placeMap["place_latitude"].(float64),
			PlaceLongitude: placeMap["place_longitude"].(float64),
			PlaceAddress:   placeMap["place_address"].(string),
		}

		tripPlaces = append(tripPlaces, place)
	}

	return tripPlaces, nil
}

// GetTripByID returns the trip by its ID
// Note: trip is a group chat
func (m *firestoreDBRepo) GetTripByID(id int64) (models.Trip, error) {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	var trip models.Trip
	var tripPlaces []models.Place

	docName := strconv.FormatInt(id, 10)

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
			PlaceTitle:     placeMap["place_title"].(string),
			PlaceLatitude:  placeMap["place_latitude"].(float64),
			PlaceLongitude: placeMap["place_longitude"].(float64),
			PlaceAddress:   placeMap["place_address"].(string),
		}

		tripPlaces = append(tripPlaces, place)
	}

	trip.TripID = id
	trip.TripPlaces = tripPlaces

	return trip, nil
}

// DeleteTripPlacesListByID deletes trip places list by its ID
// Note: trip is a group chat
func (m *firestoreDBRepo) DeleteTripPlacesListByID(id int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	docName := strconv.FormatInt(id, 10)

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
