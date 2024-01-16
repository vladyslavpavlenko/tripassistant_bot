package repository

import "github.com/vladyslavpavlenko/tripassistant_bot/internal/models"

type DatabaseRepo interface {
	AddUser(user models.User) error
	DeleteUserByID(userID int64) error
	CheckIfUserIsRegisteredByID(userID int64) (bool, error)
	AddTrip(trip models.Trip) error
	DeleteTripByID(tripID int64) error
	AddPlaceToListByTripID(place models.Place, tripID int64) error
	GetTripPlacesListByID(tripID int64) ([]models.Place, error)
	GetTripByID(tripID int64) (models.Trip, error)
	DeleteTripPlacesListByID(tripID int64) error
	DeleteTripPlaceByTitle(placeTitle string, tripID int64) error
	GetAllUserIDs() ([]int64, error)
}
