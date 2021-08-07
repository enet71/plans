package commands

import (
	"fmt"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func SendMainKeyboard(bot *tgbotapi.BotAPI) {
	me, _ := bot.GetMe()
	msg := tgbotapi.NewMessageToChannel(me.UserName, "start")
	fmt.Println(msg)
	bot.Send(msg)
}
