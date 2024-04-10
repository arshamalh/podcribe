package handlers

import (
	"podcribe/log"
	sessionPkg "podcribe/session"
	"podcribe/telegram/btns"
	"podcribe/telegram/msgs"
	"regexp"

	"go.uber.org/zap"
	"gopkg.in/telebot.v3"
)

type callbackHandler func(ctx SharedContext) error

type callbacksHandler struct {
	handlers       map[string]callbackHandler
	defaultHandler callbackHandler
	session        sessionPkg.TelegramSession
}

func NewCallbacksHandler(session sessionPkg.TelegramSession) *callbacksHandler {
	return &callbacksHandler{
		session:  session,
		handlers: make(map[string]callbackHandler),
		defaultHandler: func(ctx SharedContext) error {
			return ctx.Respond(msgs.NoHandlerHasBeenSet)
		},
	}
}

func (cbh *callbacksHandler) Register(btn btns.BtnKey, handler callbackHandler) {
	cbh.handlers[btn.String()] = handler
}

func (cbh *callbacksHandler) Handle(ctx telebot.Context) error {
	userID := ctx.Chat().ID
	unique, data := FindDataAndUnique(ctx.Callback().Data)
	ctx.Callback().Data = data
	ctx.Callback().Unique = unique
	ssn := cbh.session.GetClient(userID)
	handler, ok := cbh.handlers[unique]
	sCtx := NewSharedContext(ctx, ssn)
	if !ok {
		log.Gl.Error(
			"No handler set for button",
			zap.String("button unique", unique),
			zap.String("button data", data),
			zap.Int64("userID", userID),
		)
		return cbh.defaultHandler(sCtx)
	}
	return handler(sCtx)
}

func FindDataAndUnique(data string) (unique string, payload string) {
	callbackRx := regexp.MustCompile(`^\f([-\w]+)(\|(.+))?$`)
	if data != "" && data[0] == '\f' {
		match := callbackRx.FindAllStringSubmatch(data, -1)
		if match != nil {
			unique, payload = match[0][1], match[0][3]
			return unique, payload
		}
	}
	return "", ""
}
