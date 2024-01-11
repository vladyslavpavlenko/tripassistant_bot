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
// Note: can be either "sender" for a private chat with the inline query sender,
// "private", "group", "supergroup", or "channel"
//func ChatType(chatType string) th.Predicate {
//	return func(u telego.Update) bool {
//		if u.Message == nil {
//			return false
//		}
//		return u.Message.Chat.Type == chatType
//	}
//}

// PrivateChat is true if the message is sent in a private chat
func PrivateChat() th.Predicate {
	return func(u telego.Update) bool {
		if u.Message == nil {
			if u.MyChatMember != nil {
				return u.MyChatMember.Chat.Type == "private"
			}

			return false
		}
		return u.Message.Chat.Type == "private"
	}
}

// GroupChat is true if the message is sent in a group chat
func GroupChat() th.Predicate {
	return func(u telego.Update) bool {
		if u.Message == nil {
			if u.MyChatMember != nil {
				return u.MyChatMember.Chat.Type == "group"
			}

			return false
		}
		return u.Message.Chat.Type == "group"
	}
}

// SuperGroupChat is true if the message is sent in a supergroup chat
func SuperGroupChat() th.Predicate {
	return func(u telego.Update) bool {
		if u.Message == nil {
			if u.MyChatMember != nil {
				return u.MyChatMember.Chat.Type == "supergroup"
			}

			return false
		}
		return u.Message.Chat.Type == "supergroup"
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

// BotBlocked is true if the bot is blocked by the user
func BotBlocked() th.Predicate {
	return func(u telego.Update) bool {
		return u.MyChatMember.NewChatMember.MemberStatus() == "kicked"
	}
}

// BotAddedToGroup is true if the bot was added to a group
func BotAddedToGroup() th.Predicate {
	return func(u telego.Update) bool {
		if u.Message == nil || u.Message.NewChatMembers == nil {
			return false
		}

		for _, newMember := range u.Message.NewChatMembers {
			if newMember.Username == "tripassistant_bot" {
				return true
			}
		}

		return false
	}
}

// BotRemovedFromGroup is true if the bot was removed from a group. This can
// happen if the group was deleted or the bot itself was removed
func BotRemovedFromGroup() th.Predicate {
	return func(u telego.Update) bool {
		if u.Message == nil {
			if u.MyChatMember != nil {
				if u.MyChatMember.Chat.Type == "supergroup" {
					username := u.MyChatMember.NewChatMember.MemberUser().Username
					if username == "tripassistant_bot" {
						return u.MyChatMember.NewChatMember.MemberStatus() == "left"
					}
				}
			}

			return false
		}

		if u.Message.LeftChatMember != nil {
			if u.Message.LeftChatMember.Username == "tripassistant_bot" {
				return true
			}
		}

		return false
	}
}

// GroupNameChanged is true if the name of the group was changed
