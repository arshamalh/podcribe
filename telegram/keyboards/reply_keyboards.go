package keyboards

import (
	"podcribe/telegram/btns"
	"podcribe/telegram/msgs"

	"gopkg.in/telebot.v3"
)

func Cancel() *telebot.ReplyMarkup {
	menu := &telebot.ReplyMarkup{ResizeKeyboard: true}
	menu.Reply(btns.Cancel.AsRow(msgs.Cancel))
	return menu
}
