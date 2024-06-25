package handlers

import (
	"context"
	"fmt"
	"podcribe/log"
	"podcribe/services/convertor"
	"podcribe/telegram/msgs"

	"github.com/sashabaranov/go-openai"
	"go.uber.org/zap"
	"gopkg.in/telebot.v3"
)

func (h *handler) VoiceHandler(ctx telebot.Context) error {
	voice := ctx.Message().Voice
	// TODO: contribute to telebot to change this function also accepts and io.Writer, instead of local filename
	err := h.bot.Download(voice.MediaFile(), voice.FileID)
	if err != nil {
		log.Gl.Error(err.Error())
		return ctx.Send("can't download file from telegram, isn't it less than 20 MB?")
	}
	mp3FileName := voice.UniqueID + ".mp3"
	convertor.New().ConvertOGGToMP3(voice.FileID, mp3FileName)
	log.Gl.Info("some voice received", zap.String("name", voice.FileID))
	fmt.Println(voice.MIME)
	req := openai.AudioRequest{
		Model:    openai.Whisper1,
		FilePath: mp3FileName,
	}
	resp, err := h.openAIClient.CreateTranscription(context.Background(), req)
	if err != nil {
		fmt.Printf("Transcription error: %v\n", err)
		return ctx.Send("something unexpected happened when transcribing")
	}
	fmt.Println(resp.Text)
	return ctx.Send(msgs.FmtBasics(resp.Text))
}
