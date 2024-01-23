package handlers

import (
	"cloud.google.com/go/firestore"
	"fmt"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/config"
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/handlers/helpers"
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/mapsapi"
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/mapsapi/googleapirepo"
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/models"
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/repository"
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/repository/dbrepo"
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/responses"
	"googlemaps.github.io/maps"
	"log"
	"math/rand"
	"strings"
	"time"
)

// Repository is the Repository type
type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
	API mapsapi.APIRepo
}

// Repo the Repository used by the handlers
var Repo *Repository

// NewRepo creates a new Repository
func NewRepo(a *config.AppConfig, fsClient *firestore.Client, gmClient *maps.Client) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewFirestoreRepo(fsClient, a),
		API: googleapirepo.NewGoogleAPIRepo(gmClient, a),
	}
}

// NewTestRepo creates a new Repository
func NewTestRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewTestingFirestoreRepo(a),
		// API: googleapirepo.NewTestingGoogleAPIRepo(a),
	}
}

// NewHandlers sets the Repository for the handlers
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
	}

	if !registered {
		err := m.DB.AddUser(user)
		if err != nil {
			log.Println(err)
		}
	}

	params := &telego.SendMessageParams{
		ChatID:    tu.ID(update.Message.Chat.ID),
		Text:      responses.StartResponse,
		ParseMode: "HTML",
	}

	if update.Message.Chat.Type == "group" || update.Message.Chat.Type == "supergroup" {
		params.ReplyToMessageID = update.Message.MessageID
	}

	_, _ = bot.SendMessage(params)
}

// HelpCommandHandler handles the /help command
func (m *Repository) HelpCommandHandler(bot *telego.Bot, update telego.Update) {
	params := &telego.SendMessageParams{
		ChatID:                tu.ID(update.Message.Chat.ID),
		Text:                  responses.HelpResponse,
		ParseMode:             "HTML",
		DisableWebPagePreview: true,
	}

	if update.Message.Chat.Type == "group" || update.Message.Chat.Type == "supergroup" {
		params.ReplyToMessageID = update.Message.MessageID
	}

	_, _ = bot.SendMessage(params)
}

// AddPlaceCommandHandler handles the /addplace command
func (m *Repository) AddPlaceCommandHandler(bot *telego.Bot, update telego.Update) {
	var place models.Place

	if update.Message.ReplyToMessage.Text != "" {
		if len(update.Message.ReplyToMessage.Text) > 50 {
			params := &telego.SendMessageParams{
				ChatID:           tu.ID(update.Message.Chat.ID),
				ReplyToMessageID: update.Message.MessageID,
				Text:             responses.MessageTooLong,
				ParseMode:        "HTML",
			}

			_, _ = bot.SendMessage(params)

			return
		}

		messageText := update.Message.ReplyToMessage.Text

		var err error
		place, err = m.API.GetPlace(messageText, update.Message.Chat.Title)
		if err != nil {
			log.Println(fmt.Errorf("error parsing place from Places API: %s", err))
			place = models.Place{
				PlaceTitle: messageText,
			}
		}
	} else {
		if update.Message.ReplyToMessage.Venue != nil {
			place.PlaceTitle = update.Message.ReplyToMessage.Venue.Title
			place.PlaceLongitude = update.Message.ReplyToMessage.Venue.Location.Longitude
			place.PlaceLatitude = update.Message.ReplyToMessage.Venue.Location.Latitude
			place.PlaceAddress = update.Message.ReplyToMessage.Venue.Address
			place.PlaceID = update.Message.ReplyToMessage.Venue.GooglePlaceID

		} else if update.Message.ReplyToMessage.Location != nil {
			place.PlaceTitle = "Custom Location"
			place.PlaceLongitude = update.Message.ReplyToMessage.Location.Longitude
			place.PlaceLatitude = update.Message.ReplyToMessage.Location.Latitude

		} else {
			params := &telego.SendMessageParams{
				ChatID:           tu.ID(update.Message.Chat.ID),
				ReplyToMessageID: update.Message.MessageID,
				Text:             responses.MessageFormatNotSupported,
				ParseMode:        "HTML",
			}

			_, _ = bot.SendMessage(params)
			return
		}
	}

	err := m.DB.AddPlaceToListByTripID(place, update.Message.Chat.ID)
	if err != nil {
		fmt.Println(err)
		helpers.Error(bot, update)
		return
	}

	params := &telego.SendStickerParams{
		ChatID:           tu.ID(update.Message.Chat.ID),
		ReplyToMessageID: update.Message.ReplyToMessage.MessageID,
		Sticker: telego.InputFile{
			FileID: responses.OkSticker,
		},
	}

	_, _ = bot.SendSticker(params)
}

// RemovePlaceCommandHandler handles the /removeplace command
func (m *Repository) RemovePlaceCommandHandler(bot *telego.Bot, update telego.Update) {
	placeTitle := update.Message.ReplyToMessage.Text
	tripID := update.Message.Chat.ID

	err := m.DB.DeleteTripPlaceByTitle(placeTitle, tripID)
	if err != nil {
		fmt.Println(err)
		helpers.Error(bot, update)
		return
	}

	params := &telego.SendMessageParams{
		ChatID:           tu.ID(update.Message.Chat.ID),
		ReplyToMessageID: update.Message.MessageID,
		Text:             fmt.Sprintf(placeTitle + " was removed from the list"),
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

	if update.Message.Chat.Type == "group" || update.Message.Chat.Type == "supergroup" {
		params.ReplyToMessageID = update.Message.MessageID
	}

	_, _ = bot.SendMessage(params)
}

// AdminCommandMisuseHandler handles misuse of admin commands
func (m *Repository) AdminCommandMisuseHandler(bot *telego.Bot, update telego.Update) {
	params := &telego.SendMessageParams{
		ChatID:    tu.ID(update.Message.Chat.ID),
		Text:      responses.UseAsReplyAdmin,
		ParseMode: "HTML",
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
	tripPlaces, err := m.DB.GetTripPlacesListByID(update.Message.Chat.ID)
	if err != nil {
		log.Println(err)
	}

	var list string

	if len(tripPlaces) > 0 {
		sb := strings.Builder{}

		for i, place := range tripPlaces {
			if place.PlaceID == "" {
				sb.WriteString(fmt.Sprintf("%d. %s\n", i+1, place.PlaceTitle))
			} else {
				sb.WriteString(fmt.Sprintf("%d. <a href=\""+
					"google.com/maps/search/?api=1&query=Google&query_place_id=%s\">%s</a> ✨\n",
					i+1, place.PlaceID, place.PlaceTitle))
			}
		}

		list = sb.String()
	} else {
		list = responses.EmptyList
	}

	params := &telego.SendMessageParams{
		ChatID:                tu.ID(update.Message.Chat.ID),
		Text:                  fmt.Sprintf("<b>%s</b>\n\n%s", update.Message.Chat.Title, list),
		ParseMode:             "HTML",
		DisableWebPagePreview: true,
	}

	_, _ = bot.SendMessage(params)
}

// ClearListCommandHandler handles the /clearlist command
func (m *Repository) ClearListCommandHandler(bot *telego.Bot, update telego.Update) {
	err := m.DB.DeleteTripPlacesListByID(update.Message.Chat.ID)
	if err != nil {
		fmt.Println(err)
		helpers.Error(bot, update)
		return
	}

	params := &telego.SendStickerParams{
		ChatID:           tu.ID(update.Message.Chat.ID),
		ReplyToMessageID: update.Message.MessageID,
		Sticker: telego.InputFile{
			FileID: responses.ClearListSticker,
		},
	}

	_, _ = bot.SendSticker(params)
}

// RandomPlaceCommandHandler handles the /randomplace command
func (m *Repository) RandomPlaceCommandHandler(bot *telego.Bot, update telego.Update) {
	tripPlaces, err := m.DB.GetTripPlacesListByID(update.Message.Chat.ID)
	if err != nil {
		fmt.Println(err)
		helpers.Error(bot, update)
		return
	}

	if len(tripPlaces) == 0 {
		params := &telego.SendMessageParams{
			ChatID:           tu.ID(update.Message.Chat.ID),
			ReplyToMessageID: update.Message.MessageID,
			Text:             responses.EmptyList,
			ParseMode:        "HTML",
		}

		_, _ = bot.SendMessage(params)

		return
	}

	rand.New(rand.NewSource(time.Now().Unix()))

	randomPlace := tripPlaces[rand.Intn(len(tripPlaces))]

	if randomPlace.PlaceLatitude == 0 && randomPlace.PlaceLongitude == 0 && randomPlace.PlaceAddress == "" {
		params := &telego.SendMessageParams{
			ChatID:    tu.ID(update.Message.Chat.ID),
			Text:      randomPlace.PlaceTitle,
			ParseMode: "HTML",
		}

		_, _ = bot.SendMessage(params)
		return
	} else {
		params := &telego.SendVenueParams{
			ChatID:    tu.ID(update.Message.Chat.ID),
			Title:     randomPlace.PlaceTitle,
			Latitude:  randomPlace.PlaceLatitude,
			Longitude: randomPlace.PlaceLongitude,
			Address:   randomPlace.PlaceAddress,
		}

		_, _ = bot.SendVenue(params)
		return
	}
}

// SendPostMessageHandler sends a post message
func (m *Repository) SendPostMessageHandler(bot *telego.Bot, update telego.Update) {
	editedMessage := &telego.EditMessageTextParams{
		ChatID:    tu.ID(update.CallbackQuery.Message.Chat.ID),
		MessageID: update.CallbackQuery.Message.MessageID,
		Text:      responses.Confirmed,
		ParseMode: "HTML",
	}
	_, _ = bot.EditMessageText(editedMessage)

	userIDs, err := m.DB.GetAllUserIDs()
	if err != nil {
		log.Println(err)
		helpers.Error(bot, update)
	}

	msgText := update.CallbackQuery.Message.ReplyToMessage.Text

	btnText, btnURL, text, err := helpers.ParsePost(msgText)
	if err != nil {
		text = msgText

		for _, id := range userIDs {
			params := &telego.SendMessageParams{
				ChatID:    tu.ID(id),
				Text:      text,
				ParseMode: "HTML",
			}
			_, _ = bot.SendMessage(params)
		}
		return
	}

	inlineKeyboard := make([][]telego.InlineKeyboardButton, 0)
	inlineKeyboard = append(inlineKeyboard, []telego.InlineKeyboardButton{})
	inlineKeyboard[0] = append(inlineKeyboard[0], telego.InlineKeyboardButton{
		Text: btnText,
		URL:  btnURL,
	})

	for _, id := range userIDs {
		params := &telego.SendMessageParams{
			ChatID:      tu.ID(id),
			Text:        text,
			ParseMode:   "HTML",
			ReplyMarkup: &telego.InlineKeyboardMarkup{InlineKeyboard: inlineKeyboard},
		}
		_, _ = bot.SendMessage(params)
	}
}

// AdminPostCommandHandler handles the /post admin command
func (m *Repository) AdminPostCommandHandler(bot *telego.Bot, update telego.Update) {
	msgText := update.Message.ReplyToMessage.Text
	btnText, btnURL, text, err := helpers.ParsePost(msgText)
	if err != nil {
		text = msgText

		params := &telego.SendMessageParams{
			ChatID:    tu.ID(update.Message.Chat.ID),
			Text:      text,
			ParseMode: "HTML",
		}

		_, _ = bot.SendMessage(params)

		helpers.ConfirmationRequest(bot, update)
		return
	}

	inlineKeyboard := make([][]telego.InlineKeyboardButton, 0)
	inlineKeyboard = append(inlineKeyboard, []telego.InlineKeyboardButton{})
	inlineKeyboard[0] = append(inlineKeyboard[0], telego.InlineKeyboardButton{
		Text: btnText,
		URL:  btnURL,
	})

	params := &telego.SendMessageParams{
		ChatID:      tu.ID(update.Message.Chat.ID),
		Text:        text,
		ParseMode:   "HTML",
		ReplyMarkup: &telego.InlineKeyboardMarkup{InlineKeyboard: inlineKeyboard},
	}

	_, _ = bot.SendMessage(params)

	helpers.ConfirmationRequest(bot, update)
}

// UnknownCommandHandler handles unknown commands
func (m *Repository) UnknownCommandHandler(bot *telego.Bot, update telego.Update) {
	params := &telego.SendMessageParams{
		ChatID:    tu.ID(update.Message.Chat.ID),
		Text:      responses.UnknownCommand,
		ParseMode: "HTML",
	}

	if update.Message.Chat.Type == "group" || update.Message.Chat.Type == "supergroup" {
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

// DatabaseDeleteUserHandler handles an update when the user needs to be deleted from the Repository
func (m *Repository) DatabaseDeleteUserHandler(bot *telego.Bot, update telego.Update) {
	err := m.DB.DeleteUserByID(update.MyChatMember.From.ID)
	if err != nil {
		log.Println(err)
	}
}

// DatabaseAddTripHandler handles an update when a trip needs to be added to the Repository
func (m *Repository) DatabaseAddTripHandler(bot *telego.Bot, update telego.Update) {
	var chatID int64
	var chatTitle string

	if update.Message != nil {
		chatID = update.Message.Chat.ID
		chatTitle = update.Message.Chat.Title
	} else {
		if update.MyChatMember != nil {
			chatID = update.MyChatMember.Chat.ID
			chatTitle = update.MyChatMember.Chat.Title
		} else {
			log.Println("error adding trip to the database")
			helpers.Error(bot, update)
			return
		}
	}

	trip := models.Trip{
		TripID:     chatID,
		TripPlaces: []models.Place{},
	}

	err := m.DB.AddTrip(trip)
	if err != nil {
		log.Println(err)
		helpers.Error(bot, update)
		return
	}

	params := &telego.SendMessageParams{
		ChatID:    tu.ID(chatID),
		Text:      fmt.Sprintf(responses.NewTrip, chatTitle),
		ParseMode: "HTML",
	}

	_, _ = bot.SendMessage(params)
}

// DatabaseDeleteTripHandler handles an update when the trip needs to be deleted from the Repository
func (m *Repository) DatabaseDeleteTripHandler(bot *telego.Bot, update telego.Update) {
	if update.Message != nil {
		err := m.DB.DeleteTripByID(update.Message.Chat.ID)
		if err != nil {
			log.Println(err)
			return
		}
	}

	if update.MyChatMember != nil {
		err := m.DB.DeleteTripByID(update.MyChatMember.Chat.ID)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
