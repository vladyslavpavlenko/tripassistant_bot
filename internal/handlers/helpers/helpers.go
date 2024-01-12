package helpers

import (
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/responses"
)

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
