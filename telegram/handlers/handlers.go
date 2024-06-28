package handlers

import (
	"podcribe/repo/sqlite"
	"podcribe/session"
	"podcribe/telegram/msgs"

	"github.com/sashabaranov/go-openai"
	"gopkg.in/telebot.v3"
)

type handler struct {
	db           *sqlite.Sqlite
	bot          *telebot.Bot
	session      session.TelegramSession
	openAIClient *openai.Client
}

func New(bot *telebot.Bot, db *sqlite.Sqlite, session session.TelegramSession) *handler {
	return &handler{
		db:      db,
		bot:     bot,
		session: session,
	}
}

func (h *handler) WithOpenAIClient(client *openai.Client) {
	h.openAIClient = client
}

func (h *handler) Register() {
	h.bot.Handle("/start", func(ctx telebot.Context) error {
		session := h.session.GetClient(ctx.Chat().ID)
		return h.Start(NewSharedContext(ctx, session))
	})

	// *** Text Handlers *** //
	textHandler := NewTextHandler(h.session)
	textHandler.SetDefaultHandler(h.Default)

	textHandler.RegisterReplyKeyboardHandler(msgs.Cancel, h.Cancel)

	// *** Callback Handlers *** //
	cbsHandler := NewCallbacksHandler(h.session)

	// Handle all the texts
	h.bot.Handle(telebot.OnText, textHandler.Handle)

	// Handle all the buttons
	h.bot.Handle(telebot.OnCallback, cbsHandler.Handle)

	h.bot.Handle(telebot.OnAudio, h.AudioHandler)
	h.bot.Handle(telebot.OnVoice, h.VoiceHandler)
}
