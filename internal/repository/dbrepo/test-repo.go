package dbrepo

import (
	"errors"
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/models"
)

// AddUser adds user to the users library
func (m *testDBRepo) AddUser(user models.User) error {
	return nil
}

// DeleteUserByID adds user to the users library
func (m *testDBRepo) DeleteUserByID(id int64) error {
	return nil
}

// CheckIfUserIsRegisteredByID checks whether a user is already registered in the users library by their ID
func (m *testDBRepo) CheckIfUserIsRegisteredByID(id int64) (bool, error) {
	return false, errors.New("some error")
}

// AddTrip adds a trip to the trips collection
// Note: trip is a group chat
func (m *testDBRepo) AddTrip(trip models.Trip) error {
	return nil
}

// DeleteTripByID deletes a trip from the trips collection based on its ID
// Note: trip is a group chat
func (m *testDBRepo) DeleteTripByID(id int64) error {
	return nil
}

// UpdateTripTitle updates the title of the trip
//func (m *testDBRepo) UpdateTripTitle(trip models.Trip) error {
//	return nil
//}

// AddPlaceToListByTripID adds the place to the list of the trip identified by its ID
// Note: trip is a group chat
func (m *testDBRepo) AddPlaceToListByTripID(place models.Place, id int64) error {
	return nil
}

// GetTripPlacesListByID returns all the places from the trip by its ID
// Note: trip is a group chat
func (m *testDBRepo) GetTripPlacesListByID(id int64) ([]models.Place, error) {
	var tripPlaces []models.Place

	return tripPlaces, nil
}

// GetTripByID returns the trip by its ID
// Note: trip is a group chat
func (m *testDBRepo) GetTripByID(id int64) (models.Trip, error) {
	var trip models.Trip

	return trip, nil
}

// DeleteTripPlacesListByID deletes trip places list by its ID
// Note: trip is a group chat
func (m *testDBRepo) DeleteTripPlacesListByID(id int64) error {
	return nil
}
