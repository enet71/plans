package create_plan

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"main/bot/requests"
	"main/consts"
	"main/plan"
	"strings"
	"time"
)

func StartPlanCreating(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	messageText := createPlanText("Name", "15:04 02.01.2006", "15:04 02.01.2006", "")
	msg := tgbotapi.NewMessage(message.Chat.ID, messageText)
	msg.ReplyMarkup = createPlanMarkup()
	bot.Send(msg)
}

func FinishPlanCreating(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	plansList := plan.GetPlansList(int(message.Chat.ID))
	name, startDate, endDate, description := parseMessageText(message.Text)
	startTime, _ := time.Parse(consts.DATE_FORMAT, startDate)
	endTime, _ := time.Parse(consts.DATE_FORMAT, endDate)
	newPlan := plan.CreatePlan(name, startTime, endTime, description)
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

func updateNewPlanText(bot *tgbotapi.BotAPI, message *tgbotapi.Message, request *requests.Request, msgText string) {
	msg := tgbotapi.NewEditMessageText(message.Chat.ID, request.MessageId, msgText)
	msg.ReplyMarkup = createPlanMarkup()
	bot.Send(msg)
}

func parseMessageText(text string) (name string, startDate string, endDate string, description string) {
	fields := strings.Split(text, "\n")
	name = strings.Trim(strings.SplitN(fields[0], ":", 2)[1], " ")
	startDate = strings.Trim(strings.SplitN(fields[1], ":", 2)[1], " ")
	endDate = strings.Trim(strings.SplitN(fields[2], ":", 2)[1], " ")
	description = strings.Trim(strings.SplitN(fields[3], ":", 2)[1], " ")
	return
}

func createPlanText(name string, startDate string, endDate string, description string) string {
	return "Name: " + name + "\n" +
		"Start date: " + startDate + "\n" +
		"End date: " + endDate + "\n" +
		"Description: " + description + "\n"
}
