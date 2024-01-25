package main

import (
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	"github.com/redis/go-redis/v9"
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/handlers"
	pd "github.com/vladyslavpavlenko/tripassistant_bot/internal/handlers/predicates"
	"time"
)

const (
	throttleMaxRequests = 3
	throttleTimeWindow  = 6 * time.Second
)

// registerUpdates configures the bot handlers and middleware for different chat contexts and commands.
// Note: handlers will match only once and in order of registration.
func registerUpdates(bh *th.BotHandler, redisClient *redis.Client) {
	// Middleware
	groupChat := bh.Group(th.Or(pd.GroupChat(), pd.SuperGroupChat()))
	admin := bh.Group(pd.Admin(&app), pd.PrivateChat())

	// Globally used middleware
	bh.Use(Throttling(redisClient, throttleMaxRequests, throttleTimeWindow))

	// Global commands
	bh.Handle(handlers.Repo.StartCommandHandler, th.And(th.CommandEqual("start"), pd.PrivateChat()))
	bh.Handle(handlers.Repo.HelpCommandHandler, th.CommandEqual("help"), th.Or(pd.PrivateChat(),
		th.TextContains(app.BotUsername)))

	groupChat.Handle(handlers.Repo.AddPlaceCommandHandler, th.And(th.CommandEqual("addplace"), pd.Reply()))
	groupChat.Handle(handlers.Repo.RemovePlaceCommandHandler, th.And(th.CommandEqual("removeplace"), pd.Reply()))
	groupChat.Handle(handlers.Repo.RandomPlaceCommandHandler, th.And(th.CommandEqual("randomplace")))
	groupChat.Handle(handlers.Repo.ShowListCommandHandler, th.And(th.CommandEqual("showlist")))
	groupChat.Handle(handlers.Repo.ClearListCommandHandler, th.And(th.CommandEqual("clearlist")))

	groupChat.Handle(handlers.Repo.CommandMisuseHandler,
		th.Or(
			th.CommandEqual("addplace"),
			th.CommandEqual("removeplace"),
		),
	)

	// Ignore /start in group chats
	groupChat.Handle(func(bot *telego.Bot, update telego.Update) {}, th.CommandEqual("start"))

	bh.Handle(handlers.Repo.CommandWrongChatHandler,
		th.Or(
			th.CommandEqual("addplace"),
			th.CommandEqual("removeplace"),
			th.CommandEqual("randomplace"),
			th.CommandEqual("showlist"),
			th.CommandEqual("clearlist"),
		),
	)

	// Admin commands
	admin.Handle(handlers.Repo.AdminPostCommandHandler, th.And(th.CommandEqual("post"), pd.Reply()))

	admin.Handle(handlers.Repo.AdminCommandMisuseHandler,
		th.Or(
			th.CommandEqual("post"),
		),
	)

	// Not recognized commands
	bh.Handle(handlers.Repo.UnknownCommandHandler, th.And(th.AnyCommand(), th.TextContains(app.BotUsername)))

	// Not commands
	bh.Handle(handlers.Repo.AnyMessageHandler, th.And(th.AnyMessage(), pd.PrivateChat(), th.Not(pd.Admin(&app))))

	// Database
	bh.Handle(handlers.Repo.DatabaseDeleteUserHandler, th.And(pd.PrivateChat(), pd.BotBlocked()))
	groupChat.Handle(handlers.Repo.DatabaseAddTripHandler, pd.BotAddedToGroup(&app))
	groupChat.Handle(handlers.Repo.DatabaseDeleteTripHandler, pd.BotRemovedFromGroup(&app))

	// Callbacks
	bh.Handle(handlers.Repo.SendPostMessageHandler, th.CallbackDataContains("confirmation_pressed"))
}
