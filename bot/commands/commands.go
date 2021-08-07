package commands

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"main/bot/commands/create-plan"
	get_plans "main/bot/commands/get-plans"
	"main/bot/commands/start"
	"main/bot/requests"
)

func RunCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	command := message.Command()

	switch command {
	case Start.Strting():
		start.StartUser(bot, message)
	case CreatePlan.Strting():
		create_plan.StartPlanCreating(bot, message)
	case GetPlans.Strting():
		get_plans.GetPlans(bot, message)
	}
}

func RunCallback(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	callback := update.CallbackQuery
	switch update.CallbackQuery.Data {
	case requests.SetPlanName.String():
		create_plan.AskNewPlanName(bot, callback.Message)
		addRequest(callback)
	case requests.SetPlanStartDate.String():
		create_plan.AskNewPlanStartDate(bot, callback.Message)
		addRequest(callback)
	case requests.SetPlanEndDate.String():
		create_plan.AskNewPlanEndDate(bot, callback.Message)
		addRequest(callback)
	case requests.SetPlanDescription.String():
		create_plan.AskNewPlanDescription(bot, callback.Message)
		addRequest(callback)
	case requests.FinishPlanCreating.String():
		create_plan.FinishPlanCreating(bot, callback.Message)
	}
}

func RunRequest(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	request := requests.GetRequest(update.Message.From.ID)

	switch request.Type {
	case requests.SetPlanName.String():
		create_plan.SetNewPlanName(bot, update.Message, request)
	case requests.SetPlanStartDate.String():
		create_plan.SetNewPlanStartDate(bot, update.Message, request)
	case requests.SetPlanEndDate.String():
		create_plan.SetNewPlanEndDate(bot, update.Message, request)
	case requests.SetPlanDescription.String():
		create_plan.SetNewPlanDescription(bot, update.Message, request)
	}
}

func addRequest(callback *tgbotapi.CallbackQuery) {
	request := requests.Request{
		Type:        callback.Data,
		UserId:      callback.From.ID,
		MessageId:   callback.Message.MessageID,
		MessageText: callback.Message.Text,
	}
	requests.AddRequest(&request)
}
