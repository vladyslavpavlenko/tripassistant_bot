package dbrepo

import (
	"context"
	"errors"
	"fmt"
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/models"
	"google.golang.org/api/iterator"
	"log"
	"time"
)

// AddUser adds user to the users library
func (m *firestoreDBRepo) AddUser(user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, _, err := m.Client.Collection("users").Add(ctx, map[string]any{
		"user_id":   user.UserID,
		"user_name": user.UserName,
	})
	if err != nil {
		log.Println(fmt.Errorf("failed adding a user to the users library: %v", err))
		return err
	}

	return nil
}

// CheckIfUserIsRegisteredByID checks whether a user is already registered in the users library by their ID
func (m *firestoreDBRepo) CheckIfUserIsRegisteredByID(id int64) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	usersCollection := m.Client.Collection("users")

	query := usersCollection.Where("user_id", "==", id)
	iter := query.Documents(ctx)

	for {
		_, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			return false, nil
		}

		if err != nil {
			log.Println(fmt.Errorf("failed searching for the user in users libary: %v", err))
			return false, err
		}

		return true, nil
	}
}
