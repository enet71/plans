package get_plans

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"main/plan"
)

func GetPlans(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	plans := plan.GetPlansList(message.From.ID).Plans

	keyboardRows := make([][]tgbotapi.InlineKeyboardButton, len(plans))

	for i, plan := range plans {
		row := tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(plan.Name, plan.Uuid),
		)

		keyboardRows[i] = row
	}

	plansKeyboard := tgbotapi.NewInlineKeyboardMarkup(keyboardRows...)

	msg := tgbotapi.NewMessage(message.Chat.ID, "Test")
	msg.ReplyMarkup = plansKeyboard
	bot.Send(msg)
}
