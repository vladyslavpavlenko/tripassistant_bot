package main

import (
	th "github.com/mymmrac/telego/telegohandler"
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/config"
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/handlers"
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/handlers/predicates"
)

// Note: handlers will match only once and in order of registration
func registerHandlers(bh *th.BotHandler, app *config.AppConfig) {
	// Global commands
	bh.Handle(handlers.StartCommandHandler, th.CommandEqual("start"))

	// Admin commands
	bh.Handle(handlers.AdminPostHandler, th.And(
		th.CommandEqual("post"),
		predicates.Admin(app)))

	// Not recognized commands
	bh.Handle(handlers.UnknownCommandHandler, th.AnyCommand())

	// Not commands
	bh.Handle(handlers.AnyMessageHandler, th.Any())
}
