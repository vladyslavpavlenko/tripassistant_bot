package helpers

import (
	"fmt"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/responses"
	"regexp"
	"strings"
)

// ConfirmationRequest prompts the user to confirm some action
func ConfirmationRequest(bot *telego.Bot, update telego.Update) {
	inlineKeyboard := make([][]telego.InlineKeyboardButton, 0)
	inlineKeyboard = append(inlineKeyboard, []telego.InlineKeyboardButton{})
	inlineKeyboard[0] = append(inlineKeyboard[0], telego.InlineKeyboardButton{
		Text:         "Confirm",
		CallbackData: "confirmation_pressed",
	})

	params := &telego.SendMessageParams{
		ChatID:           tu.ID(update.Message.Chat.ID),
		Text:             responses.Confirm,
		ReplyMarkup:      &telego.InlineKeyboardMarkup{InlineKeyboard: inlineKeyboard},
		ReplyToMessageID: update.Message.ReplyToMessage.MessageID,
	}

	_, _ = bot.SendMessage(params)
}

// ServerError notifies the user about an error on the server side
func ServerError(bot *telego.Bot, update telego.Update) {
	messageParams := &telego.SendMessageParams{
		ChatID:    tu.ID(update.Message.Chat.ID),
		Text:      responses.ServerError,
		ParseMode: "HTML",
	}

	_, _ = bot.SendMessage(messageParams)

	stickerParams := &telego.SendStickerParams{
		ChatID: tu.ID(update.Message.Chat.ID),
		Sticker: telego.InputFile{
			FileID: responses.ErrorSticker,
		},
	}

	_, _ = bot.SendSticker(stickerParams)
}

// ParsePost parses button text, button URL, and post text from a string
func ParsePost(input string) (string, string, string, error) {
	pattern := `b: "(.+)"\s+u: "(.+)"\s+t: "(.+)"`
	re := regexp.MustCompile(pattern)

	matches := re.FindStringSubmatch(input)

	if len(matches) != 4 {
		return "", "", "", fmt.Errorf("invalid input format")
	}

	buttonCaption := matches[1]
	buttonLink := matches[2]
	postText := matches[3]

	postText = strings.ReplaceAll(postText, "\\n", "\n")

	return buttonCaption, buttonLink, postText, nil
}
