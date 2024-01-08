package handlers

import (
	"fmt"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

// AnyMessageHandler handles all the message that are not commands
func AnyMessageHandler(bot *telego.Bot, update telego.Update) {
	_, _ = bot.SendMessage(tu.Message(
		tu.ID(update.Message.Chat.ID),
		fmt.Sprintf("Use commands to interact with me")))
}

// StartCommandHandler handles the /start command
func StartCommandHandler(bot *telego.Bot, update telego.Update) {
	_, _ = bot.SendMessage(tu.Message(
		tu.ID(update.Message.Chat.ID),
		fmt.Sprintf("Hello %s!", update.Message.From.FirstName),
	))
}

// UnknownCommandHandler handles unknown commands
func UnknownCommandHandler(bot *telego.Bot, update telego.Update) {
	_, _ = bot.SendMessage(tu.Message(
		tu.ID(update.Message.Chat.ID),
		fmt.Sprintf("Unknown command, use /start")))
}

// AdminPostHandler handles the /post admin command
func AdminPostHandler(bot *telego.Bot, update telego.Update) {
	_, _ = bot.SendMessage(tu.Message(
		tu.ID(update.Message.Chat.ID),
		fmt.Sprintf("Admin command")))
}
