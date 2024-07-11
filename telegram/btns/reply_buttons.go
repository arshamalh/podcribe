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

func (rb ReplyBtn) AsRow() telebot.Row {
	return telebot.Row{rb.Build()}
}

const (
	Credit         ReplyBtn = ReplyBtn(msgs.Credit)
	Cancel         ReplyBtn = ReplyBtn(msgs.Cancel)
	ReferFriends   ReplyBtn = ReplyBtn(msgs.ReferFriends)
	BotLanguage    ReplyBtn = ReplyBtn(msgs.BotLanguage)
	VoicesLanguage ReplyBtn = ReplyBtn(msgs.VoicesLanguage)
	AboutUs        ReplyBtn = ReplyBtn(msgs.AboutUs)
	VoicesList     ReplyBtn = ReplyBtn(msgs.VoicesList)
)
