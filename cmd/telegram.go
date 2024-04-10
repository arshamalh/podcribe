package cmd

import (
	"context"
	"podcribe/log"
	"podcribe/repo/sqlite"
	"podcribe/session/ephemeral"
	"podcribe/telegram/handlers"
	"sync"
	"time"

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
	handler.Register()
	log.Gl.Info("Starting telegram bot")
	bot.Start()
}
