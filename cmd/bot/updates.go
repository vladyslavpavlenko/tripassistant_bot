package main

import (
	th "github.com/mymmrac/telego/telegohandler"
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/handlers"
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/handlers/predicates"
)

// Note: handlers will match only once and in order of registration
func registerUpdates(bh *th.BotHandler, r *handlers.Repository) {
	// Middleware
	reg := bh.Group(predicates.ChatType("private"))
	reg.Use(isRegistered(r))

	// Global commands
	bh.Handle(handlers.Repo.StartCommandHandler, th.CommandEqual("start"))

	bh.Handle(handlers.Repo.HelpCommandHandler, th.CommandEqual("help"))

	bh.Handle(handlers.Repo.AddPlaceCommandHandler, th.And(th.CommandEqual("addplace"), predicates.ChatType("group"), predicates.Reply()))
	bh.Handle(handlers.Repo.RemovePlaceCommandHandler, th.And(th.CommandEqual("removeplace"), predicates.ChatType("group"), predicates.Reply()))

	bh.Handle(handlers.Repo.CommandMisuseHandler,
		th.And(
			th.Or(
				th.CommandEqual("randomplace"),
				th.CommandEqual("removeplace"),
			),
			predicates.ChatType("group"),
		),
	)

	bh.Handle(handlers.Repo.CommandWrongChatHandler,
		th.Or(
			th.CommandEqual("randomplace"),
			th.CommandEqual("removeplace"),
		),
	)

	// Admin commands
	bh.Handle(handlers.Repo.AdminPostCommandHandler, th.And(predicates.Admin(&app), th.CommandEqual("post")))

	// Not recognized commands
	bh.Handle(handlers.Repo.UnknownCommandHandler, th.AnyCommand())

	// Not commands
	bh.Handle(handlers.Repo.AnyMessageHandler, th.And(predicates.ChatType("private"), th.Any()))
}
