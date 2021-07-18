package commands

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"main/consts"
	"main/plan"
	"strconv"
	"strings"
	"time"
)

func CreatePlan(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	commandArguments := message.CommandArguments()
	plansList := plan.GetList()
	planName := "Plan" + strconv.Itoa(plansList.GetLength())
	startDate := time.Now()
	endDate := time.Now().Add(time.Minute * 30)

	if commandArguments != "" {
		setArgumentsValues(commandArguments, &planName, &startDate, &endDate)
	}

	newPlan := plan.CreatePlan(planName, startDate, endDate)
	plansList.AddPlan(newPlan)
	msg := tgbotapi.NewMessage(message.Chat.ID, "Plan was created")
	bot.Send(msg)
}

func setArgumentsValues(commandArguments string, name *string, startDate *time.Time, endDate *time.Time) error {
	var err error
	arguments := strings.Fields(commandArguments)
	*name = arguments[0]
	*startDate, err = time.Parse(consts.DATE_FORMAT, arguments[1]+" "+arguments[2])

	if err != nil {
		return err
	}

	*endDate, err = time.Parse(consts.DATE_FORMAT, arguments[3]+" "+arguments[4])

	if err != nil {
		return err
	}

	return err
}
