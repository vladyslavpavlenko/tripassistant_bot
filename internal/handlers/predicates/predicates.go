package predicates

import (
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/config"
)

func Admin(app *config.AppConfig) th.Predicate {
	return func(u telego.Update) bool {
		return u.Message.From.ID == app.AdminID
	}
}
