package commands

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"main/plan"
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
		PlansList: &plan.PlanList{},
	}

	users.AddUser(user)
}
