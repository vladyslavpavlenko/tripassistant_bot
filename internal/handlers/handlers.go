package handlers

import (
	"fmt"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/responses"
)

// AnyMessageHandler handles all the message that are not commands
func AnyMessageHandler(bot *telego.Bot, update telego.Update) {
	params := &telego.SendMessageParams{
		ChatID:    tu.ID(update.Message.Chat.ID),
		Text:      responses.UseCommands,
		ParseMode: "HTML",
	}

	_, _ = bot.SendMessage(params)
}

// StartCommandHandler handles the /start command
func StartCommandHandler(bot *telego.Bot, update telego.Update) {
	params := &telego.SendMessageParams{
		ChatID:    tu.ID(update.Message.Chat.ID),
		Text:      responses.StartResponse,
		ParseMode: "HTML",
	}

	if update.Message.Chat.Type == "group" {
		params.ReplyToMessageID = update.Message.MessageID
	}

	_, _ = bot.SendMessage(params)
}

// UnknownCommandHandler handles unknown commands
func UnknownCommandHandler(bot *telego.Bot, update telego.Update) {
	params := &telego.SendMessageParams{
		ChatID:    tu.ID(update.Message.Chat.ID),
		Text:      responses.UnknownCommand,
		ParseMode: "HTML",
	}

	if update.Message.Chat.Type == "group" {
		params.ReplyToMessageID = update.Message.MessageID
	}

	_, _ = bot.SendMessage(params)
}

// AdminPostCommandHandler handles the /post admin command
func AdminPostCommandHandler(bot *telego.Bot, update telego.Update) {
	//_, _ = bot.SendMessage(tu.Message(
	//	tu.ID(update.Message.Chat.ID),
	//	fmt.Sprintf("Admin command")))
	//_, _ = bot.SendLocation(tu.Location(update.Message.Chat.ChatID(), 49.80128398674975, 24.01616258114543))
	_, _ = bot.SendVenue(tu.Venue(
		tu.ID(update.Message.Chat.ID),
		50.444211288061624,
		30.544653525946593,
		"Реберня на Аресенальній",
		"вулиця Івана Мазепи, 1, Київ, 02000",
	))
}

// HelpCommandHandler handles the /help command
func HelpCommandHandler(bot *telego.Bot, update telego.Update) {
	params := &telego.SendMessageParams{
		ChatID:    tu.ID(update.Message.Chat.ID),
		Text:      responses.HelpResponse,
		ParseMode: "HTML",
	}

	if update.Message.Chat.Type == "group" {
		params.ReplyToMessageID = update.Message.MessageID
	}

	_, _ = bot.SendMessage(params)
}

// AddPlaceCommandHandler handles the /addplace command
func AddPlaceCommandHandler(bot *telego.Bot, update telego.Update) {
	params := &telego.SendMessageParams{
		ChatID:           tu.ID(update.Message.Chat.ID),
		ReplyToMessageID: update.Message.MessageID,
		Text:             fmt.Sprintf("/addplace TBD"),
		ParseMode:        "HTML",
	}

	_, _ = bot.SendMessage(params)
}

// RemovePlaceCommandHandler handles the /removeplace command
func RemovePlaceCommandHandler(bot *telego.Bot, update telego.Update) {
	params := &telego.SendMessageParams{
		ChatID:           tu.ID(update.Message.Chat.ID),
		ReplyToMessageID: update.Message.MessageID,
		Text:             fmt.Sprintf("/removeplace TBD"),
		ParseMode:        "HTML",
	}

	_, _ = bot.SendMessage(params)
}

// CommandMisuseHandler handles misuse of commands
func CommandMisuseHandler(bot *telego.Bot, update telego.Update) {
	params := &telego.SendMessageParams{
		ChatID:    tu.ID(update.Message.Chat.ID),
		Text:      responses.UseAsReply,
		ParseMode: "HTML",
	}

	if update.Message.Chat.Type == "group" {
		params.ReplyToMessageID = update.Message.MessageID
	}

	_, _ = bot.SendMessage(params)
}

// CommandWrongChatHandler handles cases when commands are used in a private chat instead of a group
func CommandWrongChatHandler(bot *telego.Bot, update telego.Update) {
	params := &telego.SendMessageParams{
		ChatID:    tu.ID(update.Message.Chat.ID),
		Text:      responses.UseInGroups,
		ParseMode: "HTML",
	}

	_, _ = bot.SendMessage(params)
}
