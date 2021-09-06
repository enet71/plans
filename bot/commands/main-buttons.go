package commands

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func SendMainKeyboard(bot *tgbotapi.BotAPI) {
	me, _ := bot.GetMe()
	msg := tgbotapi.NewMessageToChannel(me.UserName, "start")
	bot.Send(msg)
}
