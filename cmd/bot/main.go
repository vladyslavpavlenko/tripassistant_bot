package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/config"
	"log"
	"os"
	"strconv"
)

var app config.AppConfig

func main() {
	// Run application
	bot, bh, err := run()
	if err != nil {
		log.Fatal(err)
	}

	// Stop handling updates
	defer bh.Stop()

	// Stop getting updates
	defer bot.StopLongPolling()

	// Start handling updates
	bh.Start()
}

func run() (*telego.Bot, *th.BotHandler, error) {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error getting environment variables:", err)
	}

	// Get environment variables
	botToken := os.Getenv("TOKEN")
	adminID := os.Getenv("ADMIN_ID")

	app.AdminID, _ = strconv.ParseInt(adminID, 10, 64)

	fmt.Println("Starting bot...")

	// Create bot and enable debugging info
	bot, err := telego.NewBot(botToken, telego.WithDefaultLogger(true, true))
	if err != nil {
		return nil, nil, err
	}

	botUser, err := bot.GetMe()
	if err != nil {
		return nil, nil, err
	}

	fmt.Println(fmt.Sprintf("Bot runs on @%s", botUser.Username))

	// Get updates channel
	updates, _ := bot.UpdatesViaLongPolling(nil)

	// Create bot handler and specify from where to get updates
	bh, _ := th.NewBotHandler(bot, updates)

	// Register handlers
	registerHandlers(bh, &app)

	return bot, bh, nil
}
