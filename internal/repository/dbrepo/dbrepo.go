package dbrepo

import (
	"cloud.google.com/go/firestore"
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/config"
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/database"
)

type firestoreDBRepo struct {
	App    *config.AppConfig
	Client *firestore.Client
}

type testDBRepo struct {
	App    *config.AppConfig
	Client *firestore.Client
}

func NewFirestoreRepo(client *firestore.Client, app *config.AppConfig) database.DatabaseRepo {
	return &firestoreDBRepo{
		App:    app,
		Client: client,
	}
}

func NewTestingFirestoreRepo(app *config.AppConfig) database.DatabaseRepo {
	return &testDBRepo{
		App: app,
	}
}
