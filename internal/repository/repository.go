package repository

import "github.com/vladyslavpavlenko/tripassistant_bot/internal/models"

type DatabaseRepo interface {
	AddUser(user models.User) error
	DeleteUserByID(id int64) error
	CheckIfUserIsRegisteredByID(id int64) (bool, error)
	AddTrip(trip models.Trip) error
	DeleteTripByID(id int64) error
	UpdateTripTitle(trip models.Trip) error
}
