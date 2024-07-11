package keyboards

import (
	"podcribe/telegram/btns"
	"podcribe/telegram/msgs"

	"gopkg.in/telebot.v3"
)

// When buttons doesn't have any data attached, we make them once, it help us prevent unnecessary function calls
var (
// addCalculationBtnRow    = btns.AddCalculation.AsRow(msgs.AddCalculationBtn)
// addComplexBillBtnRow    = btns.AddComplexBill.AsRow(msgs.AddComplexBillBtn)
)

func Credit() *telebot.ReplyMarkup {
	keyboard := &telebot.ReplyMarkup{}
	keyboard.Inline(
		btns.ChargesList.AsRow(msgs.ChargesList),
		btns.CancelInline.AsRow(msgs.Cancel),
	)
	return keyboard
}
