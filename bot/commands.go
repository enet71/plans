package bot

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"main/bot/commands"
)

func RunCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	command := message.Command()

	switch command {
	case CreatePlan.Strting():
		commands.CreatePlan(bot, message)
	case GetPlans.Strting():
		commands.GetPlans(bot, message)
	}
}

type CommandType int

const (
	CreatePlan CommandType = iota
	GetPlans
)

func (commandType CommandType) Strting() string {
	return [...]string{"CreatePlan", "GetPlans"}[commandType]
}
