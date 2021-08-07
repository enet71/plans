package create_plan

import (
	"github.com/Syfaro/telegram-bot-api"
	"main/bot/requests"
)

func createPlanMarkup() *tgbotapi.InlineKeyboardMarkup {
	markup := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Set name", requests.SetPlanName.String()),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Set start date", requests.SetPlanStartDate.String()),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Set end date", requests.SetPlanEndDate.String()),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Set description", requests.SetPlanDescription.String()),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Create plan", requests.FinishPlanCreating.String()),
		),
	)

	return &markup
}
