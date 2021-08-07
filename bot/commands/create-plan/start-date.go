package create_plan

import (
	"github.com/Syfaro/telegram-bot-api"
	"main/bot/requests"
)

func AskNewPlanStartDate(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Send plan start date")
	bot.Send(msg)
}

func SetNewPlanStartDate(bot *tgbotapi.BotAPI, message *tgbotapi.Message, request *requests.Request) {
	name, _, endDate, description := parseMessageText(request.MessageText)
	msgText := createPlanText(name, message.Text, endDate, description)
	updateNewPlanText(bot, message, request, msgText)
}
