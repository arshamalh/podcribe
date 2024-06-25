package handlers

import (
	"context"
	"podcribe/log"
	"podcribe/services/convertor"
	"podcribe/telegram/msgs"

	"github.com/sashabaranov/go-openai"
	"gopkg.in/telebot.v3"
)

func (h *handler) VoiceHandler(ctx telebot.Context) error {
	voice := ctx.Message().Voice
	mp3FileName := voice.UniqueID + ".mp3"

	if err := h.bot.Download(voice.MediaFile(), voice.UniqueID); err != nil {
		log.Gl.Error(err.Error())
		return ctx.Send("can't download file from telegram, isn't it less than 20 MB?")
	}

	if err := convertor.New().ConvertOGGToMP3(voice.UniqueID, mp3FileName); err != nil {
		log.Gl.Error(err.Error())
		return ctx.Send("unable to convert telegram voice to AI understandable format")
	}
	resp, err := h.openAIClient.
		CreateTranscription(context.Background(), openai.AudioRequest{
			Model:    openai.Whisper1,
			FilePath: mp3FileName,
		})

	if err != nil {
		log.Gl.Error(err.Error())
		return ctx.Send("something unexpected happened when transcribing")
	}

	return ctx.Send(msgs.FmtBasics(resp.Text))
}
