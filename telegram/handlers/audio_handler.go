package handlers

import (
	"context"
	"errors"
	"podcribe/config"
	"podcribe/log"
	"podcribe/telegram/msgs"
	"podcribe/tools"
	"slices"

	"github.com/sashabaranov/go-openai"
	"go.uber.org/zap"
	"gopkg.in/telebot.v3"
)

func (h *handler) AudioHandler(ctx telebot.Context) error {
	audio := ctx.Message().Audio
	// TODO: contribute to telebot to change this function also accepts and io.Writer, instead of local filename
	err := h.bot.Download(audio.MediaFile(), audio.FileName)
	if err != nil {
		log.Gl.Error(err.Error())
		return ctx.Send(msgs.CantDownloadFile)
	}
	log.Gl.Info("some voice received", zap.String("name", audio.FileName))

	// TODO: IMPORTANT: Add some validations for the received audio,
	//   is it the right length? is the length shorter than remaining balance?

	if !slices.Contains(tools.SupportedFormats, audio.MIME) {
		return ctx.Send(msgs.FileTypeNotSupported + config.Get().AdminUsername)
	}

	// // *** Audio stuff *** //
	// TODO: contribute to openai by making NewAudioRequest function which can be a builder pattern and limit user inputs
	req := openai.AudioRequest{
		Model:    openai.Whisper1,
		FilePath: audio.FileName,
	}

	resp, err := h.openAIClient.CreateTranscription(context.Background(), req)
	if err != nil {
		log.Gl.Error(errors.Join(err).Error())
		return ctx.Send(msgs.CantTranscribe)
	}

	return ctx.Send(msgs.FmtBasics(resp.Text))
}
