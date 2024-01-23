package main

import (
	"cloud.google.com/go/firestore"
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/config"
	"log"
)

var app config.AppConfig

func main() {
	// Run application
	bot, bh, firestoreClient, err := run()
	if err != nil {
		log.Fatal(err)
	}

	// Close Firebase connection
	defer func(firestoreClient *firestore.Client) {
		err := firestoreClient.Close()
		if err != nil {
			log.Printf("Error closing Firestore client: %v", err)
		}
	}(firestoreClient)

	// Stop handling updates
	defer bh.Stop()

	// Stop getting updates
	defer bot.StopLongPolling()

	// Start handling updates
	bh.Start()
}
