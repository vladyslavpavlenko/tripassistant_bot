package main

import (
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/config"
	"log"
)

var app config.AppConfig

func main() {
	// Run application
	bot, bh, client, err := run()
	if err != nil {
		log.Fatal(err)
	}

	// Close Firebase connection
	defer client.Close()

	// Stop handling updates
	defer bh.Stop()

	// Stop getting updates
	defer bot.StopLongPolling()

	// Start handling updates
	bh.Start()
}
