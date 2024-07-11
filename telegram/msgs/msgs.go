package msgs

// Messages
const (
	WelcomeMessage = `
Welcome to our bot, you can simply send a voice message, 
or an audio message, we will transcribe it!
`

	NoHandlerHasBeenSetMsg = `No handlers has been set`

	CreditMsg = `
Your current balance is: {balance} $
We currently support crypto payments through TON, Tron, and USDT (trc20),
Send any amount you want to charge to one of these wallets:
TON address: ''{ton}''
Tron address: ''{tron}''
USDT (trc20): ''{usdt}''
Then send your transaction ID for us.
`
)

// Buttons text
const (
	// Reply
	Cancel         string = "Cancel 🚫"
	Credit         string = "Credit 💰"
	ReferFriends   string = "Refer Friends 💌"
	BotLanguage    string = "Bot Language 🇺🇸" // TODO: show the right flag when the language is not English
	VoicesLanguage string = "Voices Language 🇺🇸"
	AboutUs        string = "About Us 🔮"
	VoicesList     string = "Voices List 📄"

	// Inline
	ChargesList string = "Transactions List 📊"
)
