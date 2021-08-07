package create_plan

import (
	"github.com/Syfaro/telegram-bot-api"
	"main/bot/requests"
)

func AskNewPlanName(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Send plan name")
	bot.Send(msg)
}

func SetNewPlanName(bot *tgbotapi.BotAPI, message *tgbotapi.Message, request *requests.Request) {
	_, startDate, endDate, description := parseMessageText(request.MessageText)
	msgText := createPlanText(message.Text, startDate, endDate, description)
	updateNewPlanText(bot, message, request, msgText)
}
