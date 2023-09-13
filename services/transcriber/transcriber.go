package transcriber

import (
	"os"

	whispergo "github.com/ggerganov/whisper.cpp/bindings/go/pkg/whisper"
)

type I interface {
	Transcribe(filepath string) (transcription string, err error)
}

type transcriber struct {
}

func New() *transcriber {
	return &transcriber{}
}

func (t transcriber) Transcribe(filepath string) (transcription string, err error) {
	modelpath := "aimodels/ggml-base.en.bin" // Path to the model, TODO: read this value from settings.yaml

	// Load the model
	model, err := whispergo.New(modelpath)
	if err != nil {
		return "", err
	}
	defer model.Close()

	// TODO: will change hard-coded transcription.srt later
	transcriptionpath := "transcription.srt"
	transcription_file, err := os.OpenFile(transcriptionpath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return "", err
	}
	defer transcription_file.Close()

	err = Process(model, filepath, os.Stdout, transcription_file)
	if err != nil {
		return "", err
	}

	return transcriptionpath, nil

}
