package handlers

import (
	"fmt"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/responses"
)

// AnyMessageHandler handles all the message that are not commands
func AnyMessageHandler(bot *telego.Bot, update telego.Update) {
	_, _ = bot.SendMessage(tu.Message(
		tu.ID(update.Message.Chat.ID),
		responses.UseCommands))
}

// StartCommandHandler handles the /start command
func StartCommandHandler(bot *telego.Bot, update telego.Update) {
	_, _ = bot.SendMessage(tu.Message(
		tu.ID(update.Message.Chat.ID),
		responses.StartResponse,
	))
}

// UnknownCommandHandler handles unknown commands
func UnknownCommandHandler(bot *telego.Bot, update telego.Update) {
	_, _ = bot.SendMessage(tu.Message(
		tu.ID(update.Message.Chat.ID),
		responses.UnknownCommand))
}

// AdminPostCommandHandler handles the /post admin command
func AdminPostCommandHandler(bot *telego.Bot, update telego.Update) {
	_, _ = bot.SendMessage(tu.Message(
		tu.ID(update.Message.Chat.ID),
		fmt.Sprintf("Admin command")))
}

// HelpCommandHandler handles the /help command
func HelpCommandHandler(bot *telego.Bot, update telego.Update) {
	_, _ = bot.SendMessage(tu.Message(
		tu.ID(update.Message.Chat.ID),
		responses.HelpResponse))
}
