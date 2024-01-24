package main

import (
	"context"
	"fmt"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	"github.com/redis/go-redis/v9"
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/handlers"
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/handlers/helpers"
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/models"
	"log"
	"strings"
	"time"
)

// Throttling creates a telego middleware to limit the number of requests a user can make in a given time window.
// It uses a Redis client to track the number of requests from each user. If a user exceeds the maxRequests limit
// within the timeWindow, a throttle message is sent and further updates are not processed by other handlers.
func Throttling(redisClient *redis.Client, maxRequests int64, timeWindow time.Duration) th.Middleware {
	return func(bot *telego.Bot, update telego.Update, next th.Handler) {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		if update.Message != nil && update.Message.From != nil && isCommand(update.Message.Text) {
			userID := update.Message.From.ID

			redisKey := fmt.Sprintf("throttle:%d", userID)

			requestCount, err := redisClient.Incr(ctx, redisKey).Result()
			if err != nil {
				log.Printf("error incrementing request count for user %d: %v", userID, err)
				next(bot, update)
				return
			}

			if requestCount == 1 {
				redisClient.Expire(ctx, redisKey, timeWindow)
			}

			if requestCount > maxRequests {
				log.Printf("user %d exceeded maximum request limit. Throttling...", userID)
				helpers.ThrottlingMessage(bot, update)
			} else {
				next(bot, update)
			}
		} else {
			next(bot, update)
		}
	}
}

// IsRegistered creates a telego middleware to check if a user is registered in the user library.
// If a user is not registered, it adds the user to the database.
func IsRegistered(m *handlers.Repository) th.Middleware {
	return func(bot *telego.Bot, update telego.Update, next th.Handler) {
		fmt.Println(fmt.Sprintf("%sUsing isRegistered middleware%s ", "\u001B[31m", "\u001B[0m"))

		registered, err := m.DB.CheckIfUserIsRegisteredByID(update.Message.From.ID)
		if err != nil {
			log.Println(err)
			return
		}

		if !registered {
			user := models.User{
				UserID:   update.Message.From.ID,
				UserName: update.Message.From.Username,
			}

			err := m.DB.AddUser(user)
			if err != nil {
				log.Println(err)
				return
			}
		}

		next(bot, update)
	}
}

// isCommand checks if the message is a command.
func isCommand(messageText string) bool {
	return strings.HasPrefix(messageText, "/")
}
