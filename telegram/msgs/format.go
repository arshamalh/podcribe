package msgs

import (
	"fmt"
	"podcribe/config"
	"strings"
)

// characters ()_-.=+><! are reserved by telegram, so we should escape them.
func FmtBasics(input string) string {
	return strings.NewReplacer(
		"''", "`", // Format mono
		"(", "\\(",
		")", "\\)",
		"_", "\\_",
		".", "\\.",
		"-", "\\-",
		"=", "\\=",
		">", "\\>",
		"<", "\\<",
		"+", "\\+",
		"!", "\\!",
	).Replace(input)
}

func FmtWelcome(name string) string {
	response := strings.NewReplacer(
		"{name}", name,
	).Replace(WelcomeMessage)
	return FmtBasics(response)
}

func FmtCredit(balance float64) string {
	wallets := config.Get().Wallets
	response := strings.NewReplacer(
		"{balance}", fmt.Sprintf("%.3f", balance),
		"{ton}", wallets.TON,
		"{tron}", wallets.TRON,
		"{usdt}", wallets.USDT_TRC20,
	).Replace(CreditMsg)
	return FmtBasics(response)
}

func FmtColonNumber(text string, number float64) string {
	return fmt.Sprintf("%s: %.2f", text, number)
}
