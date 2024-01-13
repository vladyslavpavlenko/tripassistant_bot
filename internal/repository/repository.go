package repository

import "github.com/vladyslavpavlenko/tripassistant_bot/internal/models"

type DatabaseRepo interface {
	AddUser(user models.User) error
	DeleteUserByID(id int64) error
	CheckIfUserIsRegisteredByID(id int64) (bool, error)
	AddTrip(trip models.Trip) error
	DeleteTripByID(id int64) error
	AddPlaceToListByTripID(place models.Place, id int64) error
	GetTripPlacesListByID(id int64) ([]models.Place, error)
	GetTripByID(id int64) (models.Trip, error)
	DeleteTripPlacesListByID(id int64) error
}
