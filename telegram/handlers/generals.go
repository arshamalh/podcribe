package handlers

import (
	"context"
	"podcribe/entities"
	"podcribe/log"
	"podcribe/telegram/msgs"

	"gopkg.in/telebot.v3"
)

func (h *handler) Start(ctx SharedContext) error {
	userChatID := ctx.Chat().ID
	// TODO: Add first name field to users table and add a profile button for future use

	if _, err := h.db.GetUserByChatID(context.TODO(), userChatID); err != nil {
		log.Gl.Error(err.Error())
		if err := h.db.AddUser(context.TODO(), &entities.User{
			ChatID: userChatID,
		}); err != nil {
			// User doesn't exist neither created! seek up the problem
			log.Gl.Error(err.Error())
		}
	}

	return ctx.Send(
		msgs.FmtWelcome(ctx.Message().Sender.FirstName),
	)
}

// General Text handler sits behind any other text handler,
// it decides where the text belongs according to the current scene
func (h *handler) Default(ctx SharedContext) error {
	ctx.Send("You are lost! contact bot maintainers.")
	return h.Start(ctx)
}

// We SHOULD respond to button click events,
// but when we don't want to,
// we should at least call the empty Respond so the button flash light (UI of waiting) would go away.
func (h *handler) EmptyResponder(ctx telebot.Context) {
	if err := ctx.Respond(); err != nil {
		log.Gl.Error(err.Error())
	}
}
