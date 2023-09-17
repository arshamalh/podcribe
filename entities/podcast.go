package entities

import (
	"fmt"
	"path"
	"podcribe/repo"
	"time"
)

type Podcast struct {
	Id                int       `json:"id"`
	Name              string    `json:"name"`
	PageLink          string    `json:"page_link"`
	Mp3Link           string    `json:"mp3_link"`
	Provider          string    `json:"provider"`
	Mp3Path           string    `json:"mp3_path"`
	WavPath           string    `json:"wav_path"`
	TranscriptionPath string    `json:"transcription_path"`
	TranslationPath   string    `json:"translation_path"`
	ReferencedCount   int       `json:"referenced_count"`
	CreatedAt         time.Time `json:"created_at"`
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

func (p Podcast) Model(fields map[string]interface{}) (model repo.Podcast) {
	model = repo.Podcast{
		PageLink:          p.PageLink,
		Mp3Link:           p.Mp3Link,
		Provider:          p.Provider,
		Mp3Path:           p.Mp3Path,
		WavPath:           p.WavPath,
		TranscriptionPath: p.TranscriptionPath,
		TranslationPath:   p.TranslationPath,
		ReferencedCount:   p.ReferencedCount,
		CreatedAt:         p.CreatedAt,
	}

	for key, value := range fields {
		switch key {
		case "id":
			if id, ok := value.(int); ok {
				model.Id = id
			}
		case "page_link":
			if pageLink, ok := value.(string); ok {
				model.PageLink = pageLink
			}
		case "mp3_link":
			if mp3Link, ok := value.(string); ok {
				model.Mp3Link = mp3Link
			}
		case "provider":
			if provider, ok := value.(string); ok {
				model.Provider = provider
			}
		case "mp3_path":
			if mp3Path, ok := value.(string); ok {
				model.Mp3Path = mp3Path
			}
		case "wav_path":
			if wavPath, ok := value.(string); ok {
				model.WavPath = wavPath
			}
		case "transcription_path":
			if transcriptionPath, ok := value.(string); ok {
				model.TranscriptionPath = transcriptionPath
			}
		case "translation_path":
			if translationPath, ok := value.(string); ok {
				model.TranslationPath = translationPath
			}
		case "referenced_count":
			if referencedCount, ok := value.(int); ok {
				model.ReferencedCount = referencedCount
			}
		case "created_at":
			if createdAt, ok := value.(time.Time); ok {
				model.CreatedAt = createdAt
			}
		}
	}

	return model
}
