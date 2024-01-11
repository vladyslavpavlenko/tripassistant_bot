package dbrepo

import (
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
		"trip_title":  trip.TripTitle,
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
