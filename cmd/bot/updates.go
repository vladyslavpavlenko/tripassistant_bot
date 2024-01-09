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
	bh.Handle(handlers.HelpCommandHandler, th.CommandEqual("help"))

	bh.Handle(handlers.AddPlaceCommandHandler, th.And(th.CommandEqual("addplace"), predicates.ChatType("group"), predicates.Reply()))
	bh.Handle(handlers.CommandMisuseHandler, th.And(th.CommandEqual("addplace"), predicates.ChatType("group")))
	bh.Handle(handlers.CommandWrongChatHandler, th.And(th.CommandEqual("addplace")))

	// Admin commands
	bh.Handle(handlers.AdminPostCommandHandler, th.And(predicates.Admin(app), th.CommandEqual("post")))

	// Not recognized commands
	bh.Handle(handlers.UnknownCommandHandler, th.AnyCommand())

	// Not commands
	bh.Handle(handlers.AnyMessageHandler, th.And(predicates.ChatType("private"), th.Any()))
}
