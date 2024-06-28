package handlers

import (
	"context"
	"errors"
	"podcribe/config"
	"podcribe/entities"
	"podcribe/log"
	"podcribe/services/convertor"
	"podcribe/telegram/msgs"

	"github.com/sashabaranov/go-openai"
	"go.uber.org/zap"
	"gopkg.in/telebot.v3"
)

func (h *handler) VoiceHandler(ctx telebot.Context) error {
	voice := ctx.Message().Voice
	mp3FileName := voice.UniqueID + ".mp3"
	userID := ctx.Chat().ID
	flowContext := context.TODO()
	invoiceFactor := config.Get().InvoiceFactor

	minimumBalance := invoiceFactor * float64(voice.Duration)

	user, err := h.db.GetUserByChatID(flowContext, userID)
	if err != nil || user.Balance < minimumBalance {
		errNotEnoughBalance := msgs.NotEnoughBalance(minimumBalance)
		log.Gl.Error(
			errors.Join(err, errNotEnoughBalance).Error(),
			zap.Int64("userChatID", userID),
		)
		return ctx.Send(msgs.FmtBasics(errNotEnoughBalance.Error()))
		// TODO: Add a keyboard to charge right away
		// TODO: TODO: change the message in a way that required-current
	}

	if err := h.bot.Download(voice.MediaFile(), voice.UniqueID); err != nil {
		log.Gl.Error(err.Error())
		return ctx.Send(msgs.CantDownloadFile)
	}

	if err := convertor.ConvertOGGToMP3(voice.UniqueID, mp3FileName); err != nil {
		log.Gl.Error(err.Error())
		return ctx.Send(msgs.UnableToConvert)
	}

	// New audio entity
	voiceModel := entities.NewAudio(user.ID, voice.FilePath, voice.UniqueID, ctx.Message().ID)
	if err := h.db.AddAudio(flowContext, voiceModel); err != nil {
		log.Gl.Error(
			err.Error(),
			zap.String("uniqueID", voice.UniqueID),
			zap.String("@", "VoiceHandler@AddAudio"),
		)
		// Don't break the user flow because we were unable to add new audios
	}

	resp, err := h.openAIClient.
		CreateTranscription(context.Background(), openai.AudioRequest{
			Model:    openai.Whisper1,
			FilePath: mp3FileName,
		})

	if err != nil {
		log.Gl.Error(err.Error())
		return ctx.Send(msgs.CantTranscribe)
	}

	log.Gl.Info(
		"OpenAI audio duration",
		zap.Int64("audioID", voiceModel.ID),
		zap.Float64("duration", resp.Duration),
	)

	// resp.Duration seems to be empty 🤔
	voiceModel.Duration = voice.Duration
	voiceModel.Transcription = resp.Text
	if err := h.db.UpdateAudio(flowContext, voiceModel); err != nil {
		log.Gl.Error(
			err.Error(),
			zap.String("uniqueID", voice.UniqueID),
			zap.String("@", "VoiceHandler@UpdateAudio"),
		)
		// Don't break the user flow because we were unable to add new audios
	}
	// TODO:
	// resp.GetRateLimitHeaders().RemainingRequests
	// Buffer new requests, and define and update global rate limiter.

	// TODO: Check if AudioUniqueID was repetitive, don't transcribe it again (but calculate the price)

	// TODO: voice duration can't be wrong, so we can trust the value got from API,
	//   but audio duration is easy to manipulate, so needs validation
	invoice := voiceModel.IssueInvoice(invoiceFactor)
	if err := h.db.AddInvoice(flowContext, invoice); err != nil {
		err = errors.Join(errors.New("unable to update balance and make invoice"), err)
		log.Gl.Error(
			err.Error(),
			zap.Int64("invoice.AudioID", invoice.AudioID),
			zap.String("@", "VoiceHandler@AddInvoice"),
		)
	}

	return ctx.Send(msgs.FmtBasics(resp.Text))
}
