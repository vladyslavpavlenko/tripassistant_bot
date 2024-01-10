package repository

import "github.com/vladyslavpavlenko/tripassistant_bot/internal/models"

type DatabaseRepo interface {
	AddUser(user models.User) error
	CheckIfUserIsRegisteredByID(id int64) (bool, error)
}
