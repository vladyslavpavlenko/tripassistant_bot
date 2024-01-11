package predicates

import (
	"github.com/mymmrac/telego"
	"github.com/stretchr/testify/assert"
	"github.com/vladyslavpavlenko/tripassistant_bot/internal/config"
	"testing"
)

func TestAdmin(t *testing.T) {
	tests := []struct {
		name           string
		update         telego.Update
		expectedResult bool
	}{
		{
			name:           "nil_message",
			update:         telego.Update{},
			expectedResult: false,
		},

		{
			name:           "admin_update",
			update:         telego.Update{Message: &telego.Message{From: &telego.User{ID: 123}}},
			expectedResult: true,
		},

		{
			name:           "non-admin_update",
			update:         telego.Update{Message: &telego.Message{From: &telego.User{ID: 999}}},
			expectedResult: false,
		},
	}

	mockApp := &config.AppConfig{
		AdminIDs: []int64{123, 456},
	}

	for _, e := range tests {
		t.Run(e.name, func(t *testing.T) {
			predicate := Admin(mockApp)
			assert.Equal(t, e.expectedResult, predicate(e.update))
		})
	}
}

func TestPrivateChat(t *testing.T) {
	tests := []struct {
		name           string
		update         telego.Update
		expectedResult bool
	}{
		{
			name:           "nil_message",
			update:         telego.Update{},
			expectedResult: false,
		},

		{
			name:           "private_private",
			update:         telego.Update{Message: &telego.Message{Chat: telego.Chat{Type: "private"}}},
			expectedResult: true,
		},

		{
			name:           "private_sender",
			update:         telego.Update{Message: &telego.Message{Chat: telego.Chat{Type: "sender"}}},
			expectedResult: false,
		},

		{
			name:           "private_group",
			update:         telego.Update{Message: &telego.Message{Chat: telego.Chat{Type: "group"}}},
			expectedResult: false,
		},

		{
			name:           "private_supergroup",
			update:         telego.Update{Message: &telego.Message{Chat: telego.Chat{Type: "supergroup"}}},
			expectedResult: false,
		},

		{
			name:           "private_channel",
			update:         telego.Update{Message: &telego.Message{Chat: telego.Chat{Type: "channel"}}},
			expectedResult: false,
		},
	}

	for _, e := range tests {
		t.Run(e.name, func(t *testing.T) {
			predicate := PrivateChat()
			assert.Equal(t, e.expectedResult, predicate(e.update))
		})
	}
}

func TestGroupChat(t *testing.T) {
	tests := []struct {
		name           string
		update         telego.Update
		expectedResult bool
	}{
		{
			name:           "nil_message",
			update:         telego.Update{},
			expectedResult: false,
		},

		{
			name:           "group_group",
			update:         telego.Update{Message: &telego.Message{Chat: telego.Chat{Type: "group"}}},
			expectedResult: true,
		},

		{
			name:           "group_sender",
			update:         telego.Update{Message: &telego.Message{Chat: telego.Chat{Type: "sender"}}},
			expectedResult: false,
		},

		{
			name:           "group_private",
			update:         telego.Update{Message: &telego.Message{Chat: telego.Chat{Type: "private"}}},
			expectedResult: false,
		},

		{
			name:           "group_supergroup",
			update:         telego.Update{Message: &telego.Message{Chat: telego.Chat{Type: "supergroup"}}},
			expectedResult: false,
		},

		{
			name:           "group_channel",
			update:         telego.Update{Message: &telego.Message{Chat: telego.Chat{Type: "channel"}}},
			expectedResult: false,
		},
	}

	for _, e := range tests {
		t.Run(e.name, func(t *testing.T) {
			predicate := GroupChat()
			assert.Equal(t, e.expectedResult, predicate(e.update))
		})
	}
}

func TestSuperGroupChat(t *testing.T) {
	tests := []struct {
		name           string
		update         telego.Update
		expectedResult bool
	}{
		{
			name:           "nil_message",
			update:         telego.Update{},
			expectedResult: false,
		},

		{
			name:           "supergroup_supergroup",
			update:         telego.Update{Message: &telego.Message{Chat: telego.Chat{Type: "supergroup"}}},
			expectedResult: true,
		},

		{
			name:           "supergroup_sender",
			update:         telego.Update{Message: &telego.Message{Chat: telego.Chat{Type: "sender"}}},
			expectedResult: false,
		},

		{
			name:           "supergroup_private",
			update:         telego.Update{Message: &telego.Message{Chat: telego.Chat{Type: "private"}}},
			expectedResult: false,
		},

		{
			name:           "supergroup_group",
			update:         telego.Update{Message: &telego.Message{Chat: telego.Chat{Type: "group"}}},
			expectedResult: false,
		},

		{
			name:           "supergroup_channel",
			update:         telego.Update{Message: &telego.Message{Chat: telego.Chat{Type: "channel"}}},
			expectedResult: false,
		},
	}

	for _, e := range tests {
		t.Run(e.name, func(t *testing.T) {
			predicate := SuperGroupChat()
			assert.Equal(t, e.expectedResult, predicate(e.update))
		})
	}
}

func TestReply(t *testing.T) {
	tests := []struct {
		name           string
		update         telego.Update
		expectedResult bool
	}{
		{
			name:           "nil_message",
			update:         telego.Update{},
			expectedResult: false,
		},

		{
			name:           "reply",
			update:         telego.Update{Message: &telego.Message{ReplyToMessage: &telego.Message{}}},
			expectedResult: true,
		},

		{
			name:           "non-reply",
			update:         telego.Update{Message: &telego.Message{}},
			expectedResult: false,
		},
	}

	for _, e := range tests {
		t.Run(e.name, func(t *testing.T) {
			predicate := Reply()
			assert.Equal(t, e.expectedResult, predicate(e.update))
		})
	}
}
