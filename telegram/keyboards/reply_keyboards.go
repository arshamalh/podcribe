package keyboards

import (
	"podcribe/telegram/btns"

	"gopkg.in/telebot.v3"
)

func Main() *telebot.ReplyMarkup {
	menu := &telebot.ReplyMarkup{ResizeKeyboard: true}
	menu.Reply(
		menu.Row(
			btns.Credit.Build(),
			btns.ReferFriends.Build(),
		),
		menu.Row(
			btns.VoicesLanguage.Build(),
			btns.VoicesList.Build(),
		),
		menu.Row(
			btns.BotLanguage.Build(),
			btns.AboutUs.Build(),
		),
	)
	return menu
}

func Cancel() *telebot.ReplyMarkup {
	menu := &telebot.ReplyMarkup{ResizeKeyboard: true}
	menu.Reply(btns.Cancel.AsRow())
	return menu
}
