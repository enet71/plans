package bot

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"main/bot/commands"
)

func RunCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	command := message.Command()

	switch command {
	case Start.Strting():
		commands.StartUser(bot, message)
	case CreatePlan.Strting():
		commands.CreatePlan(bot, message)
	case GetPlans.Strting():
		commands.GetPlans(bot, message)
	}
}

type CommandType int

const (
	Start CommandType = iota
	CreatePlan
	GetPlans
)

func (commandType CommandType) Strting() string {
	return [...]string{"start", "createPlan", "getPlans"}[commandType]
}
