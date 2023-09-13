package entities

import (
	"fmt"
	"path"
)

type Podcast struct {
	Name              string
	PageLink          string
	Mp3Link           string
	Mp3Path           string
	WavPath           string
	TranscriptionPath string
	TranslationPath   string
}

// TODO: Add an ID for file paths and don't just relay on p.Name
func (p Podcast) GetTranscriptionPath(root string) string {
	p.TranscriptionPath = path.Join(root, fmt.Sprintf("%s.srt", p.Name))
	return p.TranscriptionPath
}

func (p Podcast) GetWavPath(root string) string {
	p.WavPath = path.Join(root, fmt.Sprintf("%s.wav", p.Name))
	return p.WavPath
}
