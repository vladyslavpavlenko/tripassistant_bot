package dbrepo

import (
	"cloud.google.com/go/firestore"
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/config"
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/repository"
)

type firestoreDBRepo struct {
	App    *config.AppConfig
	Client *firestore.Client
}

type testDBRepo struct {
	App    *config.AppConfig
	Client *firestore.Client
}

// NewFirestoreRepo creates a new instance of a repository.DatabaseRepo that uses Firestore as the backend database.
func NewFirestoreRepo(client *firestore.Client, app *config.AppConfig) repository.DatabaseRepo {
	return &firestoreDBRepo{
		App:    app,
		Client: client,
	}
}

// NewTestingFirestoreRepo creates a new instance of a repository.DatabaseRepo intended for testing purposes.
func NewTestingFirestoreRepo(app *config.AppConfig) repository.DatabaseRepo {
	return &testDBRepo{
		App: app,
	}
}
