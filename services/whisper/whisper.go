package whisper

type I interface {
	Transcribe(filepath string) (transcription string, err error)
}

type whisper struct {
}

func New() *whisper {
	return &whisper{}
}

func (w whisper) Transcribe(filepath string) (transcription string, err error) {
	return "", nil
}
