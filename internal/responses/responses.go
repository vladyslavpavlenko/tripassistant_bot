package responses

const (
	StartResponse  = "<b>Hello! 👋</b>\n\nI will help you create lists of places you want to visit during your trips.\n\nTo begin, add me to the group where you're planning your trip and use commands to interact with me.\n\nLet's get adventurous!"
	UnknownCommand = "Unknown command, use /help"
	HelpResponse   = "You can control me by sending these commands:\n\n/addplace (as a reply to a message) adds a new place to the list.\n\n<b>For simple lists:</b>\nLviv National Opera\n\n<b>For venue lists:</b>\nLviv National Opera\n49.844127370749945, 24.026233754747057\nSvobody Ave, 28, Lviv, Lviv Oblast, 79000\n\nYou can also reply to the <a href=\"https://telegram.org/blog/captions-places#places\">venue message</a> with this command to add it to the venue list.\n\n/removeplace (as a reply to a message) removes a place from the list. \n\nReply to the message that contains the name of the place previously added to the list or to the venue message.\n\n/randomplace returns a random place from the list.\n\n/showlist returns the whole list.\n\n/clearlist clears the list contents."
	UseCommands    = "Use commands to interact with me"
	UseAsReply     = "Respond to a message with this command. For more details, use /help"
	UseInGroups    = "Use this command in a group chat. For more details, use /help"
	EmptyList      = "There are no places in this trip yet. For more details, use /help"
	MessageTooLong = "This message is too long"
	ServerError    = "<b>Oops! Something went wrong... ⚙️</b>\n\nTry re-adding the bot to the group or restarting the dialogue. If the issue persists, contact the developer"
	OkSticker      = "CAACAgIAAxkBAAJtM2WhkXY7isE6_ku8vLa3y3Ke-OSjAAL-AANWnb0K2gRhMC751_80BA"
	ErrorSticker   = "CAACAgIAAxkBAAJtNWWhmiB5sSEBWxm9us9SNCHHHTqLAAICAQACVp29Ck7ibIHLQOT_NAQ"
)
