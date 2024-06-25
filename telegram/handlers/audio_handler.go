package handlers

import (
	"context"
	"fmt"
	"podcribe/log"
	"podcribe/telegram/msgs"

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
		return ctx.Send("can't download file from telegram, isn't it less than 20 MB?")
	}
	log.Gl.Info("some voice received", zap.String("name", audio.FileName))

	// // *** Audio stuff *** //
	req := openai.AudioRequest{
		Model:    openai.Whisper1,
		FilePath: audio.FileName,
	}
	resp, err := h.openAIClient.CreateTranscription(context.Background(), req)
	if err != nil {
		fmt.Printf("Transcription error: %v\n", err)
		return ctx.Send("something unexpected happened when transcribing")
	}
	fmt.Println(resp.Text)
	return ctx.Send(msgs.FmtBasics(resp.Text))
}
