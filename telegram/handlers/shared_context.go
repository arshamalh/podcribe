package handlers

import (
	"podcribe/session"

	"gopkg.in/telebot.v3"
)

type SharedContext interface {
	telebot.Context
	session.UserDataSession
}

type sharedContext struct {
	telebot.Context
	session.UserDataSession
}

func NewSharedContext(context telebot.Context, session session.UserDataSession) sharedContext {
	return sharedContext{context, session}
}
