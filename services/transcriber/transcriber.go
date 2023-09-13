package transcriber

import (
	"os"
	"podcribe/entities"

	whispergo "github.com/ggerganov/whisper.cpp/bindings/go/pkg/whisper"
)

type I interface {
	Transcribe(podcast *entities.Podcast) error
}

type transcriber struct {
	// TODO: IMPORTANT Should replace with a storage Adapter having some methods
	// That adapter should be able to accept local files, or Something like minio or S3,
	// And return io.Writers
	rootStorage string
}

func New() *transcriber {
	return &transcriber{}
}

func (t transcriber) Transcribe(podcast *entities.Podcast) error {
	modelpath := "aimodels/ggml-base.en.bin" // Path to the model, TODO: read this value from settings.yaml

	// Load the model
	model, err := whispergo.New(modelpath)
	if err != nil {
		return err
	}
	defer model.Close()

	// TODO: will change hard-coded transcription.srt later
	transcription_file, err := os.OpenFile(podcast.GetTranscriptionPath(t.rootStorage), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer transcription_file.Close()

	return Process(model, podcast.WavPath, os.Stdout, transcription_file)

}
