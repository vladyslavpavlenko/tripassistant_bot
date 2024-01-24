package main

import (
	"cloud.google.com/go/firestore"
	"github.com/redis/go-redis/v9"
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/config"
	"log"
)

var app config.AppConfig

func main() {
	// Run application
	services, err := setup()
	if err != nil {
		log.Fatal(err)
	}

	// Close Firebase connection
	defer func(client *firestore.Client) {
		err := client.Close()
		if err != nil {
			log.Printf("error closing Firestore client: %v", err)
		}
	}(services.FirestoreClient)

	// Close Redis connection
	defer func(client *redis.Client) {
		err := client.Close()
		if err != nil {
			log.Printf("error closing Redis client: %v", err)
		}
	}(services.RedisClient)

	// Stop handling updates
	defer services.BotHandler.Stop()

	// Stop getting updates
	defer services.Bot.StopLongPolling()

	// Start handling updates
	services.BotHandler.Start()
}
