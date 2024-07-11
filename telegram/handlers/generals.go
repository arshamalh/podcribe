package handlers

import (
	"context"
	"podcribe/config"
	"podcribe/entities"
	"podcribe/log"
	"podcribe/telegram/keyboards"
	"podcribe/telegram/msgs"

	"gopkg.in/telebot.v3"
)

func (h *handler) Start(ctx SharedContext) error {
	userChatID := ctx.Chat().ID

	user, err := h.db.GetUserByChatID(context.TODO(), userChatID)
	if user != nil && err == nil {
		return ctx.Send(
			msgs.FmtWelcome(user.TFName),
			keyboards.Main(),
		)
	}

	log.Gl.Error(err.Error())
	user = &entities.User{
		TFName: ctx.Sender().FirstName,
		TLName: ctx.Sender().LastName,
		ChatID: userChatID,
	}
	if err := h.db.AddUser(context.TODO(), user); err != nil {
		// User doesn't exist neither created! seek up the problem
		log.Gl.Error(err.Error())
		message := "your account can't be created! message the admin: " + config.Get().AdminUsername
		return ctx.Send(msgs.FmtBasics(message), keyboards.Main())
	}

	return ctx.Send(
		msgs.FmtWelcome(user.TFName),
		keyboards.Main(),
	)
}

// General Text handler sits behind any other text handler,
// it decides where the text belongs according to the current scene
func (h *handler) Default(ctx SharedContext) error {
	ctx.Send("You are lost! message the admin: " + config.Get().AdminUsername)
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
