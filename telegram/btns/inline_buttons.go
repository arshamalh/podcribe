package btns

import (
	"strings"

	"gopkg.in/telebot.v3"
)

type BtnKey string

// Turns button unique to a pair of (`\f` + button string value) so it can accessible for handlers.
func (bk BtnKey) Key() string {
	return "\f" + string(bk)
}

// Turns button unique directly to its string value without any extra information attached.
func (bk BtnKey) String() string {
	return string(bk)
}

func (bk BtnKey) Build(text string, data ...string) telebot.Btn {
	return telebot.Btn{
		Unique: string(bk),
		Text:   text,
		Data:   strings.Join(data, "|"),
	}
}

func (bk BtnKey) AsRow(text string, data ...string) telebot.Row {
	return telebot.Row{bk.Build(text, data...)}
}

const (
	AddCalculation BtnKey = "calcAdd"
)
