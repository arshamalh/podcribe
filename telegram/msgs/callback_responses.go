package msgs

import "gopkg.in/telebot.v3"

var (
	NoCalculations      = NewCallbackResponse(NoCalculationsMsg)
	NoHandlerHasBeenSet = NewCallbackResponse(NoHandlerHasBeenSetMsg)
)

func NewCallbackResponse(text string) *telebot.CallbackResponse {
	return &telebot.CallbackResponse{Text: text}
}
