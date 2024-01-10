package main

import (
	"fmt"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/handlers"
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/models"
	"log"
)

// isRegistered checks if the user is in the user library
func isRegistered(m *handlers.Repository) th.Middleware {
	return func(bot *telego.Bot, update telego.Update, next th.Handler) {
		fmt.Println(fmt.Sprintf("%sUsing isRegistered middleware%s ", "\u001B[31m", "\u001B[0m"))

		registered, err := m.DB.CheckIfUserIsRegisteredByID(update.Message.From.ID)
		if err != nil {
			log.Println(err)
			// TODO: Handle the error appropriately
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
				// TODO: Handle the error appropriately
				return
			}
		}

		next(bot, update)
	}
}
