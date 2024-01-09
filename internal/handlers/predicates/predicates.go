package predicates

import (
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/config"
	"slices"
)

// Admin is true if the message is sent by an admin
func Admin(app *config.AppConfig) th.Predicate {
	return func(u telego.Update) bool {
		if u.Message == nil {
			return false
		}
		return slices.Contains(app.AdminIDs, u.Message.From.ID)
	}
}

// ChatType is true if the message is sent in a chat of the specified type
func ChatType(chatType string) th.Predicate {
	return func(u telego.Update) bool {
		if u.Message == nil {
			return false
		}
		return u.Message.Chat.Type == chatType
	}
}

// Reply is true if the message is sent in reply to another message
func Reply() th.Predicate {
	return func(u telego.Update) bool {
		if u.Message == nil {
			return false
		}
		return u.Message.ReplyToMessage != nil
	}
}
