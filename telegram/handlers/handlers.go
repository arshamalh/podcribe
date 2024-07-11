package handlers

import (
	"podcribe/entities"
	"podcribe/repo/sqlite"
	"podcribe/services/ton"
	"podcribe/services/tron"
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
	ton          *ton.TON
	tron         *tron.Tron
}

func New(bot *telebot.Bot, db *sqlite.Sqlite, session session.TelegramSession) *handler {
	return &handler{
		db:      db,
		bot:     bot,
		session: session,
	}
}

func (h *handler) WithOpenAIClient(client *openai.Client) *handler {
	h.openAIClient = client
	return h
}

func (h *handler) WithTON(ton *ton.TON) *handler {
	h.ton = ton
	return h
}

func (h *handler) WithTRON(tron *tron.Tron) *handler {
	h.tron = tron
	return h
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
	textHandler.RegisterReplyKeyboardHandler(msgs.Credit, h.Credit)

	textHandler.Register()
	sceneCredit := NewSceneTextHandler(entities.SceneCredit)
	sceneCredit.Register(0, h.CreditTxTextHandler)

	// *** Callback Handlers *** //
	cbsHandler := NewCallbacksHandler(h.session)
	// cbsHandler.Register(btns.ChargesList, func(ctx SharedContext) error { return nil }) // TODO: Complete

	// Handle all the texts
	h.bot.Handle(telebot.OnText, textHandler.Handle)

	// Handle all the buttons
	h.bot.Handle(telebot.OnCallback, cbsHandler.Handle)
	h.bot.Handle(telebot.OnAudio, h.AudioHandler)
	h.bot.Handle(telebot.OnVoice, h.VoiceHandler)
}
