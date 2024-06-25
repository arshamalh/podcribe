package cmd

import (
	"context"
	"os"
	"podcribe/log"
	"podcribe/repo/sqlite"
	"podcribe/session/ephemeral"
	"podcribe/telegram/handlers"
	"sync"
	"time"

	"github.com/sashabaranov/go-openai"
	"gopkg.in/telebot.v3"
)

func startTelegram(token string, wg *sync.WaitGroup) {
	defer wg.Done()
	bot, err := telebot.NewBot(telebot.Settings{
		Token:     token,
		Poller:    &telebot.LongPoller{Timeout: 10 * time.Second},
		ParseMode: telebot.ModeMarkdownV2,
	})
	if err != nil {
		log.Gl.Error(err.Error())
	}

	// TODO: Read config file in any exists

	openAIToken := os.Getenv("OPENAI_TOKEN")
	config := openai.DefaultConfig(openAIToken)
	config.BaseURL = "https://api.gilas.io/v1"
	openAIClient := openai.NewClientWithConfig(config)

	// TODO: If there was a redis configuration and that was connectable, connect!
	// If there wasn't any, initialize an ephemeral session
	session := ephemeral.New()
	// redisSession := redis.New()
	db, err := sqlite.New("./data.db")
	if err != nil {
		log.Gl.Fatal("unable to initialize db")
	}
	db.Migrate(context.TODO())
	handler := handlers.New(bot, db, session)
	handler.WithOpenAIClient(openAIClient)
	handler.Register()
	log.Gl.Info("Starting telegram bot")
	bot.Start()
}
