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

func NewFirestoreRepo(fsClient *firestore.Client, a *config.AppConfig) repository.DatabaseRepo {
	return &firestoreDBRepo{
		App:    a,
		Client: fsClient,
	}
}

func NewTestingFirestoreRepo(a *config.AppConfig) repository.DatabaseRepo {
	return &testDBRepo{
		App: a,
	}
}
