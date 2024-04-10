package btns

import (
	"podcribe/telegram/msgs"

	"gopkg.in/telebot.v3"
)

type ReplyBtn string

func (rb ReplyBtn) Key() string {
	return string(rb)
}

func (rb ReplyBtn) Build() telebot.Btn {
	return telebot.Btn{
		Text: string(rb),
	}
}

func (rb ReplyBtn) AsRow(text string) telebot.Row {
	return telebot.Row{rb.Build()}
}

const (
	Cancel ReplyBtn = ReplyBtn(msgs.Cancel)
)
