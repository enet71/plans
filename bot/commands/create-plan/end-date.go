package create_plan

import (
	"github.com/Syfaro/telegram-bot-api"
	"main/bot/requests"
)

func AskNewPlanEndDate(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Send plan name")
	bot.Send(msg)
}

func SetNewPlanEndDate(bot *tgbotapi.BotAPI, message *tgbotapi.Message, request *requests.Request) {
	name, startDate, _, description := parseMessageText(request.MessageText)
	msgText := createPlanText(name, startDate, message.Text, description)
	updateNewPlanText(bot, message, request, msgText)
}
