package bot

import (
	"fmt"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"main/consts"
	"main/users"
	"time"
)

func StartBot() {
	bot, _ := tgbotapi.NewBotAPI(consts.TELEGRAM_TOKEN)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, _ := bot.GetUpdatesChan(u)

	go createScheduler()
	createUpdatesListener(bot, &updates)
}

func createScheduler() {
	for minute := range time.Tick(time.Minute) {
		fmt.Println(minute)
		users.ForEachUser(func(user users.User) {

		})
	}
}

func createUpdatesListener(bot *tgbotapi.BotAPI, updates *tgbotapi.UpdatesChannel) {
	for update := range *updates {
		fmt.Println(update)

		if update.CallbackQuery != nil {

		}

		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			RunCommand(bot, update.Message)
		}
	}
}
