package handlers

import (
	"podcribe/entities"
	"podcribe/log"
	sessionPkg "podcribe/session"
	"podcribe/telegram/msgs"

	"go.uber.org/zap"
	"gopkg.in/telebot.v3"
)

type textHandler func(ctx SharedContext) error

type textHandlers struct {
	session               sessionPkg.TelegramSession
	handlers              map[entities.Scene]*SceneTextHandler
	replyKeyboardHandlers map[string]textHandler
	defaultHandler        textHandler
}

func NewTextHandler(session sessionPkg.TelegramSession) textHandlers {
	return textHandlers{
		session:               session,
		handlers:              make(map[entities.Scene]*SceneTextHandler),
		replyKeyboardHandlers: make(map[string]textHandler),
		defaultHandler: func(ctx SharedContext) error {
			return ctx.Send(msgs.FmtBasics("No handler set! contact bot maintainers"))
		},
	}
}

// TODO: question 0 for default cases?
func (th textHandlers) Register(handlers ...*SceneTextHandler) {
	for _, sceneHandler := range handlers {
		th.handlers[sceneHandler.scene] = sceneHandler
	}
}

func (th textHandlers) RegisterReplyKeyboardHandler(unique string, handler textHandler) {
	th.replyKeyboardHandlers[unique] = handler
}

func (th *textHandlers) SetDefaultHandler(handler textHandler) {
	th.defaultHandler = handler
}

func (th textHandlers) Handle(ctx telebot.Context) error {
	userID := ctx.Chat().ID
	ssn := th.session.GetClient(userID)
	sCtx := NewSharedContext(ctx, ssn)

	unique := ctx.Message().Text
	if handler, ok := th.replyKeyboardHandlers[unique]; ok {
		return handler(sCtx)
	}

	// *** Get scene and its SceneTextHandler type
	scene := entities.Scene(0) // TODO: ssn.GetScene()
	sceneHandler, ok := th.handlers[scene]
	if !ok {
		log.Gl.Error(
			"No handler set for scene",
			zap.Int("scene", int(scene)),
			zap.Int64("userID", userID),
		)
		return th.defaultHandler(sCtx)
	}

	// *** Get Final Handler
	question := entities.Question(0) // TODO: ssn.GetCurrentQuestion()
	handler, ok := sceneHandler.handlers[question]
	if !ok {
		log.Gl.Error(
			"No handler set for question",
			zap.Int("scene", int(scene)),
			zap.Int("question", int(question)),
			zap.Int64("userID", userID),
		)
		return sceneHandler.defaultHandler(sCtx)
	}

	return handler(sCtx)
}

type SceneTextHandler struct {
	scene          entities.Scene
	handlers       map[entities.Question]textHandler
	defaultHandler textHandler
}

func NewSceneTextHandler(scene entities.Scene) *SceneTextHandler {
	return &SceneTextHandler{
		scene:    scene,
		handlers: make(map[entities.Question]textHandler),
		defaultHandler: func(ctx SharedContext) error {
			return ctx.Send(msgs.FmtBasics("No handler set, contact bot maintainers"))
		},
	}
}

func (sth SceneTextHandler) Register(question entities.Question, handler textHandler) {
	sth.handlers[question] = handler
}

func (sth *SceneTextHandler) SetDefaultHandler(handler textHandler) {
	sth.defaultHandler = handler
}
