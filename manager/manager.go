package manager

import (
	"podcribe/services/convertor"
	"podcribe/services/crawler"
	"podcribe/services/downloader"
	"podcribe/services/translator"
	"podcribe/services/whisper"
)

// We may want to just download the podcast using the bot (no translating, convertion)

// Holds different services which each has its own settings
// Make a new manager for each user
// Can hold its own settings too
type Manager struct {
	crawler     crawler.I
	downloader  downloader.I
	convertor   convertor.I
	transcriber whisper.I
	translator  translator.I
	settings    managerSettings
}

func New(crawler crawler.I, downloader downloader.I, convertor convertor.I, transcriber whisper.I, translator translator.I, flowType FlowType) *Manager {
	return &Manager{
		crawler, downloader, convertor, transcriber, translator,
		managerSettings{
			flowType: flowType,
		},
	}
}

// Start a full or partial flow of steps the bot is cable of
// it may just find and download the file and return the path or go down till translation
func (m Manager) Start(link string) (string, error) {
	podcast_link, err := m.crawler.Find(link)
	if err != nil {
		return "", err
	}
	filepath, err := m.downloader.Download(podcast_link)
	if err != nil {
		return "", err
	}

	if m.settings.flowType == FullFlow {
		return filepath, nil
	}

	podcast_path, err := m.convertor.Convert(filepath)
	if err != nil {
		return "", err
	}
	podcast_text, err := m.transcriber.Transcribe(podcast_path)
	if err != nil {
		return "", err
	}

	if m.settings.flowType == NoTranslation {
		return podcast_text, nil
	}

	translation, err := m.translator.Translate(podcast_text)
	if err != nil {
		return "", err
	}
	return translation, nil
}
