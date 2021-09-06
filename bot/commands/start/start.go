package start

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"main/bot/cosnt/keyboards"
	"main/users"
)

func StartUser(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	messageUser := *message.From
	user := users.User{
		ID:        messageUser.ID,
		FirstName: messageUser.FirstName,
		LastName:  messageUser.LastName,
		UserName:  messageUser.UserName,
		ChatID:    message.Chat.ID,
	}

	users.AddUser(&user)

	msg := tgbotapi.NewMessage(message.Chat.ID, "Bot was started")
	msg.ReplyMarkup = keyboards.MainKeyboard
	bot.Send(msg)
}
