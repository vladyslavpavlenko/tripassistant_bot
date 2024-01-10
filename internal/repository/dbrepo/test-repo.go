package dbrepo

import (
	"errors"
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/models"
)

// AddUser adds user to the users library
func (m *testDBRepo) AddUser(user models.User) error {
	return nil
}

// CheckIfUserIsRegisteredByID checks whether a user is already registered in the users library by their ID
func (m *testDBRepo) CheckIfUserIsRegisteredByID(id int64) (bool, error) {
	return false, errors.New("some error")
}
