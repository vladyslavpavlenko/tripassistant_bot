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

// appServices holds instances of services used in the application.
type appServices struct {
	Bot             *telego.Bot
	BotHandler      *th.BotHandler
	FirestoreClient *firestore.Client
	RedisClient     *redis.Client
	MapsClient      *maps.Client
}

// envVariables holds environment variables used in the application.
type envVariables struct {
	BotToken     string   `env:"BOT_TOKEN"`
	AdminIDs     []string `env:"ADMIN_IDS,split"`
	FirebasePath string   `env:"FIREBASE_PATH"`
	MapsAPIKey   string   `env:"MAPS_API_KEY"`
	Redis        *redisSettings
}

// redisSettings holds settings for connecting to Redis.
type redisSettings struct {
	Addr string `env:"REDIS_ADDR"`
	User string `env:"REDIS_USER"`
	Pass string `env:"REDIS_PASS"`
}

// setup sets up and initializes all necessary services like the bot, bot handler, Firestore client, and Redis client.
func setup() (*appServices, error) {
	// Get environment variables
	env, err := loadEvnVariables()
	if err != nil {
		return nil, err
	}

	// Create bot and enable debugging info
	bot, err := initBot(env.BotToken)
	if err != nil {
		return nil, err
	}

	// Connect to Firestore
	firestoreClient, err := connectToFirestore(env.FirebasePath)
	if err != nil {
		return nil, err
	}

	// Connect to Google Maps Platform API
	mapsClient, err := connectToGoogleMapsPlatform(env.MapsAPIKey)
	if err != nil {
		return nil, err
	}

	// Connect to Redis
	redisClient, err := connectToRedis(env.Redis.Addr, env.Redis.User, env.Redis.Pass)
	if err != nil {
		return nil, err
	}

	// Get updates channel
	updates, _ := bot.UpdatesViaLongPolling(nil)

	// Create bot handler and specify from where to get updates
	bh, _ := th.NewBotHandler(bot, updates)

	// Register handlers
	repo := handlers.NewRepo(&app, firestoreClient, mapsClient)
	handlers.NewHandlers(repo)
	registerUpdates(bh, redisClient)

	return &appServices{
		Bot:             bot,
		BotHandler:      bh,
		MapsClient:      mapsClient,
		RedisClient:     redisClient,
		FirestoreClient: firestoreClient,
	}, nil
}

func loadEvnVariables() (*envVariables, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error getting environment variables: %v", err)
	}

	// Get environment variables
	botToken := os.Getenv("BOT_TOKEN")

	firebasePath := os.Getenv("FIREBASE_PATH")

	mapsAPIKey := os.Getenv("MAPS_API_KEY")

	redisAddr := os.Getenv("REDIS_ADDR")
	redisUser := os.Getenv("REDIS_USER")
	redisPass := os.Getenv("REDIS_PASS")

	adminIDsStr := os.Getenv("ADMIN_IDS")
	adminIDsSlice := strings.Split(adminIDsStr, ",")

	for _, idStr := range adminIDsSlice {
		adminID, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			log.Printf("error parsing admin IDs: %v\n", err)
		}
		app.AdminIDs = append(app.AdminIDs, adminID)
	}

	return &envVariables{
		BotToken:     botToken,
		FirebasePath: firebasePath,
		MapsAPIKey:   mapsAPIKey,
		Redis: &redisSettings{
			Addr: redisAddr,
			User: redisUser,
			Pass: redisPass,
		},
	}, nil
}

// initBot initializes the Telegram bot with the given token.
func initBot(token string) (*telego.Bot, error) {
	log.Println("Starting bot...")

	bot, err := telego.NewBot(token, telego.WithDefaultLogger(true, true))
	if err != nil {
		return nil, fmt.Errorf("error initializing bot: %v", err)
	}

	botUser, err := bot.GetMe()
	if err != nil {
		return nil, fmt.Errorf("error getting bot user: %v", err)
	}

	log.Println(fmt.Sprintf("Bot runs on @%s", botUser.Username))

	return bot, nil
}

// connectToFirestore initializes a connection to the Firestore database using the given configuration path.
func connectToFirestore(configPath string) (*firestore.Client, error) {
	log.Println("Connecting to Firestore...")

	opt := option.WithCredentialsFile(configPath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing Firebase app: %v", err)
	}

	client, err := app.Firestore(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error initializing Firestore client: %v", err)
	}

	log.Println("Connected!")

	return client, nil
}

// connectToGoogleMapsPlatform initializes a connection to the Google Maps Platform using the given API key.
func connectToGoogleMapsPlatform(APIKey string) (*maps.Client, error) {
	log.Println("Connecting to Google Maps Platform...")
	client, err := maps.NewClient(maps.WithAPIKey(APIKey))
	if err != nil {
		return nil, fmt.Errorf("error connecting to Google Maps Platform: %v", err)
	}

	googleapirepo.NewGoogleAPIRepo(client, &app)

	log.Println("Connected!")

	return client, nil
}

// connectToRedis initializes a connection to the Redis using the given settings.
func connectToRedis(addr string, username string, password string) (*redis.Client, error) {
	log.Println("Connecting to Redis...")

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: username,
		Password: password,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("error connecting to Redis: %v", err)
	}

	log.Println("Connected!")

	return client, nil
}
