package handlers

import (
	"cloud.google.com/go/firestore"
	"fmt"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/config"
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/models"
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/repository"
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/repository/dbrepo"
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/responses"
	"log"
)

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

// Repo the repository used by the handlers
var Repo *Repository

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig, client *firestore.Client) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewFirestoreRepo(client, a),
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// StartCommandHandler handles the /start command
func (m *Repository) StartCommandHandler(bot *telego.Bot, update telego.Update) {
	user := models.User{
		UserID:   update.Message.From.ID,
		UserName: update.Message.From.Username,
	}

	registered, err := m.DB.CheckIfUserIsRegisteredByID(user.UserID)

	if err != nil {
		log.Println(err)
		// TODO: Revise
	}

	if !registered {
		err := m.DB.AddUser(user)
		if err != nil {
			log.Println(err)
			// TODO: Revise
		}
	}

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

// HelpCommandHandler handles the /help command
func (m *Repository) HelpCommandHandler(bot *telego.Bot, update telego.Update) {
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
func (m *Repository) AddPlaceCommandHandler(bot *telego.Bot, update telego.Update) {
	params := &telego.SendMessageParams{
		ChatID:           tu.ID(update.Message.Chat.ID),
		ReplyToMessageID: update.Message.MessageID,
		Text:             fmt.Sprintf("/addplace TBD"),
		ParseMode:        "HTML",
	}

	_, _ = bot.SendMessage(params)
}

// RemovePlaceCommandHandler handles the /removeplace command
func (m *Repository) RemovePlaceCommandHandler(bot *telego.Bot, update telego.Update) {
	params := &telego.SendMessageParams{
		ChatID:           tu.ID(update.Message.Chat.ID),
		ReplyToMessageID: update.Message.MessageID,
		Text:             fmt.Sprintf("/removeplace TBD"),
		ParseMode:        "HTML",
	}

	_, _ = bot.SendMessage(params)
}

// CommandMisuseHandler handles misuse of commands
func (m *Repository) CommandMisuseHandler(bot *telego.Bot, update telego.Update) {
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
func (m *Repository) CommandWrongChatHandler(bot *telego.Bot, update telego.Update) {
	params := &telego.SendMessageParams{
		ChatID:    tu.ID(update.Message.Chat.ID),
		Text:      responses.UseInGroups,
		ParseMode: "HTML",
	}

	_, _ = bot.SendMessage(params)
}

// ShowListCommandHandler handles the /showlist command
func (m *Repository) ShowListCommandHandler(bot *telego.Bot, update telego.Update) {
	params := &telego.SendMessageParams{
		ChatID:    tu.ID(update.Message.Chat.ID),
		Text:      "/showlist TBD",
		ParseMode: "HTML",
	}

	_, _ = bot.SendMessage(params)
}

// DeleteListCommandHandler handles the /deletelist command
func (m *Repository) DeleteListCommandHandler(bot *telego.Bot, update telego.Update) {
	params := &telego.SendMessageParams{
		ChatID:    tu.ID(update.Message.Chat.ID),
		Text:      "/deletelist TBD",
		ParseMode: "HTML",
	}

	_, _ = bot.SendMessage(params)
}

// RandomPlaceCommandHandler handles the /randomplace command
func (m *Repository) RandomPlaceCommandHandler(bot *telego.Bot, update telego.Update) {
	params := &telego.SendMessageParams{
		ChatID:    tu.ID(update.Message.Chat.ID),
		Text:      "/randomplace TBD",
		ParseMode: "HTML",
	}

	_, _ = bot.SendMessage(params)
}

// AdminPostCommandHandler handles the /post admin command
func (m *Repository) AdminPostCommandHandler(bot *telego.Bot, update telego.Update) {
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

// UnknownCommandHandler handles unknown commands
func (m *Repository) UnknownCommandHandler(bot *telego.Bot, update telego.Update) {
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

// AnyMessageHandler handles all the message that are not commands
func (m *Repository) AnyMessageHandler(bot *telego.Bot, update telego.Update) {
	params := &telego.SendMessageParams{
		ChatID:    tu.ID(update.Message.Chat.ID),
		Text:      responses.UseCommands,
		ParseMode: "HTML",
	}

	_, _ = bot.SendMessage(params)
}
