package transcriber

import (
	"os"

	fpkg "path/filepath"

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
	modelpath := "whisper/models/ggml-base.en.bin" // Path to the model, TODO: read this value from settings.yaml

	// Load the model
	model, err := whispergo.New(modelpath)
	if err != nil {
		return "", nil
	}
	defer model.Close()

	flags, err := NewFlags(fpkg.Base(os.Args[0]), os.Args[1:])
	Process(model, filepath, flags)

	return "", nil
}
