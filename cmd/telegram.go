package cmd

import (
	"context"
	"podcribe/config"
	"podcribe/log"
	"podcribe/repo/sqlite"
	"podcribe/services/ton"
	"podcribe/services/tron"
	"podcribe/session/ephemeral"
	"podcribe/telegram/handlers"
	"sync"
	"time"

	"github.com/sashabaranov/go-openai"
	"gopkg.in/telebot.v3"
)

func startTelegram(cfg config.Config, wg *sync.WaitGroup) {
	defer wg.Done()
	bot, err := telebot.NewBot(telebot.Settings{
		Token:     cfg.TelegramToken,
		Poller:    &telebot.LongPoller{Timeout: 10 * time.Second},
		ParseMode: telebot.ModeMarkdownV2,
	})
	if err != nil {
		log.Gl.Error(err.Error())
	}

	// TODO: Read config file in any exists

	openAIConfig := openai.DefaultConfig(cfg.OpenAIToken)
	openAIConfig.BaseURL = cfg.OpenAIBase
	openAIClient := openai.NewClientWithConfig(openAIConfig)

	// TODO: If there was a redis configuration and that was connectable, connect!
	// If there wasn't any, initialize an ephemeral session
	session := ephemeral.New()
	// redisSession := redis.New()
	db, err := sqlite.New("./data.db")
	if err != nil {
		log.Gl.Fatal("unable to initialize db")
	}
	db.Migrate(context.TODO())

	ton := ton.New(cfg.TON_BASE, cfg.TON_APIKey)
	tron := tron.New(cfg.TRON_BASE, cfg.TRON_APIKey)

	handler := handlers.New(bot, db, session).
		WithOpenAIClient(openAIClient).
		WithTON(ton).
		WithTRON(tron)

	handler.Register()
	log.Gl.Info("Starting telegram bot")
	bot.Start()
}
