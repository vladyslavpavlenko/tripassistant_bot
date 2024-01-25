package predicates

import (
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/config"
	"slices"
)

// Admin is true if the message is sent by an admin.
func Admin(app *config.AppConfig) th.Predicate {
	return func(u telego.Update) bool {
		if u.Message == nil {
			return false
		}
		return slices.Contains(app.AdminIDs, u.Message.From.ID)
	}
}

// PrivateChat is true if the message is sent in a private chat.
func PrivateChat() th.Predicate {
	return func(u telego.Update) bool {
		if u.Message == nil {
			if u.MyChatMember != nil {
				return u.MyChatMember.Chat.Type == telego.ChatTypePrivate
			}

			return false
		}
		return u.Message.Chat.Type == telego.ChatTypePrivate
	}
}

// GroupChat is true if the message is sent in a group chat.
func GroupChat() th.Predicate {
	return func(u telego.Update) bool {
		if u.Message == nil {
			if u.MyChatMember != nil {
				return u.MyChatMember.Chat.Type == telego.ChatTypeGroup
			}

			return false
		}
		return u.Message.Chat.Type == telego.ChatTypeGroup
	}
}

// SuperGroupChat is true if the message is sent in a supergroup chat.
func SuperGroupChat() th.Predicate {
	return func(u telego.Update) bool {
		if u.Message == nil {
			if u.MyChatMember != nil {
				return u.MyChatMember.Chat.Type == telego.ChatTypeSupergroup
			}

			return false
		}
		return u.Message.Chat.Type == telego.ChatTypeSupergroup
	}
}

// Reply is true if the message is sent in reply to another message.
func Reply() th.Predicate {
	return func(u telego.Update) bool {
		if u.Message != nil {
			return u.Message.ReplyToMessage != nil
		}
		return false
	}
}

// BotBlocked is true if the bot is blocked by the user.
func BotBlocked() th.Predicate {
	return func(u telego.Update) bool {
		if u.MyChatMember != nil && u.MyChatMember.NewChatMember != nil {
			return u.MyChatMember.NewChatMember.MemberStatus() == telego.MemberStatusBanned
		}
		return false
	}
}

// BotAddedToGroup is true if the bot was added to a group.
func BotAddedToGroup(app *config.AppConfig) th.Predicate {
	return func(u telego.Update) bool {
		if u.MyChatMember != nil {
			newMemberUsername := u.MyChatMember.NewChatMember.MemberUser().Username
			newMemberStatus := u.MyChatMember.NewChatMember.MemberStatus()
			oldMemberStatus := u.MyChatMember.OldChatMember.MemberStatus()

			if newMemberUsername == app.BotUsername && newMemberStatus != telego.MemberStatusLeft &&
				(oldMemberStatus == telego.MemberStatusLeft || oldMemberStatus == telego.MemberStatusBanned) {
				return true
			}
		}

		return false
	}
}

// BotRemovedFromGroup is true if the bot was removed from a group. This can
// happen if the group was deleted or the bot itself was removed.
func BotRemovedFromGroup(app *config.AppConfig) th.Predicate {
	return func(u telego.Update) bool {
		if u.Message != nil {
			if u.Message.LeftChatMember != nil {
				if u.Message.LeftChatMember.Username == app.BotUsername {
					return true
				}
			}
		}

		return false
	}
}
