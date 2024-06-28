package msgs

import (
	"fmt"
	"strings"
)

func FmtWelcome(name string) string {
	response := strings.NewReplacer(
		"{name}", name,
	).Replace(WelcomeMessage)
	return FmtBasics(response)
}

// characters ()_-.=+><! are reserved by telegram, so we should escape them.
func FmtBasics(input string) string {
	return strings.NewReplacer(
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

func FmtColonNumber(text string, number float64) string {
	return fmt.Sprintf("%s: %.2f", text, number)
}
