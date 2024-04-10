package transcriber

import (
	"podcribe/entities"
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
	return nil
}
