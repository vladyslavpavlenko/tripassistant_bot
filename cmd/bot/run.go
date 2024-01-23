package main

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	"github.com/redis/go-redis/v9"
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/handlers"
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/mapsapi/googleapirepo"
	"google.golang.org/api/option"
	"googlemaps.github.io/maps"
	"log"
	"os"
	"strconv"
	"strings"
)

func run() (*telego.Bot, *th.BotHandler, *firestore.Client, error) {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		return nil, nil, nil, fmt.Errorf("error getting environment variables: %v", err)
	}

	// Get environment variables
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	adminIDsStr := os.Getenv("ADMIN_IDS")
	fbConfigPath := os.Getenv("FIREBASE_CONFIG_PATH")
	gmAPIKey := os.Getenv("GOOGLE_MAPS_API_KEY")
	redisAddr := os.Getenv("REDIS_ADDR")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	adminIDsSlice := strings.Split(adminIDsStr, ",")

	for _, idStr := range adminIDsSlice {
		adminID, _ := strconv.ParseInt(idStr, 10, 64)
		app.AdminIDs = append(app.AdminIDs, adminID)
	}

	log.Println("Starting bot...")

	// Create bot and enable debugging info
	bot, err := telego.NewBot(botToken, telego.WithDefaultLogger(true, true))
	if err != nil {
		return nil, nil, nil, err
	}

	botUser, err := bot.GetMe()
	if err != nil {
		return nil, nil, nil, err
	}

	log.Println(fmt.Sprintf("Bot runs on @%s", botUser.Username))

	// Connect to Firebase
	log.Println("Connecting to Firebase...")

	opt := option.WithCredentialsFile(fbConfigPath)

	fbApp, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("error Firebase initializing app: %v", err)
	}

	// Create Firestore client
	fsClient, err := fbApp.Firestore(context.Background())
	if err != nil {
		return nil, nil, nil, fmt.Errorf("error initializing Firestore client: %v", err)
	}

	log.Println("Connected!")

	// Connect to Google Maps Platform API
	gmClient, err := maps.NewClient(maps.WithAPIKey(gmAPIKey))
	if err != nil {
		log.Println("Error connecting to Google Maps Platform API:", err)
	}

	googleapirepo.NewGoogleAPIRepo(gmClient, &app)

	// Connect to Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
	})

	pong, err := redisClient.Ping(context.Background()).Result()

	log.Println(pong, err)

	// Get updates channel
	updates, _ := bot.UpdatesViaLongPolling(nil)

	// Create bot handler and specify from where to get updates
	bh, _ := th.NewBotHandler(bot, updates)

	// Register handlers
	repo := handlers.NewRepo(&app, fsClient, gmClient)
	handlers.NewHandlers(repo)
	registerUpdates(bh, redisClient)

	return bot, bh, fsClient, nil
}
