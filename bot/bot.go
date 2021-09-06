package bot

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"main/bot/commands"
	"main/bot/requests"
	"main/consts"
	"main/plan"
	"main/users"
	"time"
)

func StartBot() {
	bot, _ := tgbotapi.NewBotAPI(consts.TELEGRAM_TOKEN)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, _ := bot.GetUpdatesChan(u)

	go createScheduler(bot)
	createUpdatesListener(bot, &updates)
}

func createScheduler(bot *tgbotapi.BotAPI) {
	for minute := range time.Tick(time.Second * 10) {

		_, offset := minute.Zone()
		utcDate := minute.UTC().Add(time.Second * time.Duration(offset))

		for _, plansList := range plan.GetPlansLists() {
			for _, plan := range plansList.Plans {
				if plan.StartDate.Sub(utcDate).Seconds() < 11 {
					user := users.GetUSer(plansList.UserId)
					msg := tgbotapi.NewMessage(user.ChatID, "Reminder")
					bot.Send(msg)
				}
			}
		}
	}
}

func createUpdatesListener(bot *tgbotapi.BotAPI, updates *tgbotapi.UpdatesChannel) {
	for update := range *updates {
		if update.Message != nil && update.Message.IsCommand() {
			commands.RunCommand(bot, update.Message)
		}

		if update.CallbackQuery != nil && requests.IsRequest(update.CallbackQuery.Data) {
			commands.RunCallback(bot, update)
		}

		if update.Message != nil && requests.UserHasRequest(update.Message.From.ID) {
			commands.RunRequest(bot, update)
		}
	}
}
