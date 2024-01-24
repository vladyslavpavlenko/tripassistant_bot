package responses

const (
	StartResponse             = "<b>Hello! 👋</b>\n\nI will help you create lists of places you want to visit during your trips.\n\nTo begin, add me to the group where you're planning your trip and use commands to interact with me.\n\nLet's get adventurous!\n\n<b>Contact:</b> @vlapavlenko / <a href=\"https://github.com/vladyslavpavlenko/tripassistant_bot\"><b>GitHub</b></a>"
	UnknownCommand            = "Unknown command, use /help"
	HelpResponse              = "<b>You can control me with these commands:</b>\n\n/addplace (as a reply to a message): Adds a new place to the list, like \"Lviv National Opera\". You can also reply to a <a href=\"https://telegram.org/blog/captions-places#places\">venue message</a> or a location message with this command.\n\nTrip Assistant automatically searches for the place on maps for your convenience. However, <b>automatic results might not always be accurate</b>, so for better search accuracy, write the name of the place in detail.\n\n/removeplace (as a reply to a message): Removes a place from the list. Reply to the message containing the name of the previously added place or to the venue message.\n\n/randomplace: Returns a random place from the list.\n\n/showlist: Displays the entire list.\n\n/clearlist: Clears all contents from the list.\n"
	NewTrip                   = "Adventure awaits! New trip: <b>%s</b>\n\nFor more details, use /help 🗺"
	UseCommands               = "Use commands to interact with me"
	UseAsReply                = "Respond to a message with this command. For more details, use /help"
	UseAsReplyAdmin           = "Respond to a message with this command.\n\n<b>Format:</b>\nb: \"\"\nu: \"\"\nt: \"\"\n\nUse HTML formatting."
	UseInGroups               = "Use this command in a group chat. For more details, use /help"
	EmptyList                 = "There are no places in this trip yet. For more details, use /help"
	MessageTooLong            = "This message is too long"
	MessageFormatNotSupported = "This message type is not supported. For more details, use /help"
	Confirm                   = "Would you like to confirm?"
	Confirmed                 = "<i>Confirmed!</i>"
	ServerError               = "<b>Oops! Something went wrong... ⚙️</b>\n\nTry re-adding the bot to the group or restarting the dialogue. If the issue persists, contact the developer"
	Throttling                = "Not so often!"
	OkSticker                 = "CAACAgIAAxkBAAJtM2WhkXY7isE6_ku8vLa3y3Ke-OSjAAL-AANWnb0K2gRhMC751_80BA"
	ClearListSticker          = "CAACAgIAAxkBAAJtk2WiscWgDG-csFsHFltulEd7LiWDAALMFQACdWepSdykWqEAAVfrBDQE"
	ErrorSticker              = "CAACAgIAAxkBAAJtl2WisvMQ-2YqcStg8NrTwfWb8enmAAInFQACN4qgScbxh817Y0veNAQ"
)
