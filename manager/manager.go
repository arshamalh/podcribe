package manager

import (
	"podcribe/services/convertor"
	"podcribe/services/crawler"
	"podcribe/services/downloader"
	"podcribe/services/transcriber"
	"podcribe/services/translator"
)

// We may want to just download the podcast using the bot (no translating, convertion)

// Holds different services which each has its own settings
// Make a new manager for each user
// Can hold its own settings too
type Manager struct {
	crawler     crawler.I
	downloader  downloader.I
	convertor   convertor.I
	transcriber transcriber.I
	translator  translator.I
}

func New(crawler crawler.I, downloader downloader.I, convertor convertor.I, transcriber transcriber.I, translator translator.I) *Manager {
	return &Manager{
		crawler, downloader, convertor, transcriber, translator,
	}
}

// the only difference is filtered providers
// TODO: Should we include a direct download for this function?
func (m Manager) JustDownload(link string) (filepath string, err error) {
	podcast_link, err := m.crawler.Find(link)
	if err != nil {
		return "", err
	}
	return m.downloader.Download(podcast_link)
}

// Start a full or partial flow of steps the bot is cable of
// it may just find and download the file and return the path or go down till translation
func (m Manager) FullFlow(link string) (translation string, transcription string, podcast_path string, err error) {
	podcast_link, err := m.crawler.Find(link)
	if err != nil {
		return "", "", "", err
	}

	filepath, err := m.downloader.Download(podcast_link)
	if err != nil {
		return "", "", "", err
	}

	podcast_path, err = m.convertor.Convert(filepath)
	if err != nil {
		return "", "", "", err
	}

	transcription, err = m.transcriber.Transcribe(podcast_path)
	if err != nil {
		return "", "", "", err
	}

	translation, err = m.translator.Translate(transcription)

	return translation, transcription, podcast_path, err
}

func (m Manager) FullExceptTranslation(link string) (transcription string, filepath string, err error) {
	podcast_link, err := m.crawler.Find(link)
	if err != nil {
		return "", "", err
	}

	filepath, err = m.downloader.Download(podcast_link)
	if err != nil {
		return "", "", err
	}

	podcast_path, err := m.convertor.Convert(filepath)
	if err != nil {
		return "", "", err
	}

	transcription, err = m.transcriber.Transcribe(podcast_path)

	return transcription, filepath, err
}

func (m Manager) TranscribeDownloadedMP3(path string) (string, error) {
	podcast_path, err := m.convertor.Convert(path)
	if err != nil {
		return "", err
	}

	return m.transcriber.Transcribe(podcast_path)
}

func (m Manager) TranslateDownloadedMP3(path string) (string, error) {
	podcast_path, err := m.convertor.Convert(path)
	if err != nil {
		return "", err
	}

	podcast_text, err := m.transcriber.Transcribe(podcast_path)
	if err != nil {
		return "", err
	}

	return m.translator.Translate(podcast_text)
}

func (m Manager) TranscribeDownloadedWAV(path string) (string, error) {
	return m.transcriber.Transcribe(path)
}

func (m Manager) TranslateDownloadedWAV(path string) (string, error) {
	podcast_text, err := m.transcriber.Transcribe(path)
	if err != nil {
		return "", err
	}

	return m.translator.Translate(podcast_text)
}
