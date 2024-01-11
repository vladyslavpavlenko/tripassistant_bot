package main

import (
	th "github.com/mymmrac/telego/telegohandler"
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/handlers"
	pd "github.com/vladyslavpavlenko/tripassistant_bot/internal/handlers/predicates"
)

// Note: handlers will match only once and in order of registration
func registerUpdates(bh *th.BotHandler, r *handlers.Repository) {
	// Middleware
	groupChat := bh.Group(pd.ChatType("group"))
	admin := bh.Group(pd.Admin(&app), pd.ChatType("private"))

	// Globally used middleware
	// privateChat.Use(IsRegistered(r))

	// Global commands
	bh.Handle(handlers.Repo.StartCommandHandler, th.CommandEqual("start"))
	bh.Handle(handlers.Repo.HelpCommandHandler, th.CommandEqual("help"))

	groupChat.Handle(handlers.Repo.AddPlaceCommandHandler, th.And(th.CommandEqual("addplace"), pd.Reply()))
	groupChat.Handle(handlers.Repo.RemovePlaceCommandHandler, th.And(th.CommandEqual("removeplace"), pd.Reply()))
	groupChat.Handle(handlers.Repo.RandomPlaceCommandHandler, th.And(th.CommandEqual("randomplace")))
	groupChat.Handle(handlers.Repo.ShowListCommandHandler, th.And(th.CommandEqual("showlist")))
	groupChat.Handle(handlers.Repo.DeleteListCommandHandler, th.And(th.CommandEqual("deletelist")))

	groupChat.Handle(handlers.Repo.CommandMisuseHandler,
		th.Or(
			th.CommandEqual("addplace"),
			th.CommandEqual("removeplace"),
		),
	)

	bh.Handle(handlers.Repo.CommandWrongChatHandler,
		th.Or(
			th.CommandEqual("addplace"),
			th.CommandEqual("removeplace"),
			th.CommandEqual("randomplace"),
			th.CommandEqual("showlist"),
			th.CommandEqual("deletelist"),
		),
	)

	// Admin commands
	admin.Handle(handlers.Repo.AdminPostCommandHandler, th.CommandEqual("post"))

	// Not recognized commands
	bh.Handle(handlers.Repo.UnknownCommandHandler, th.AnyCommand())

	// Not commands
	bh.Handle(handlers.Repo.AnyMessageHandler, th.Any())
}
