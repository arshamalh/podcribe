package manager

import (
	"podcribe/entities"
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
func (m Manager) JustDownload(link string) (*entities.Podcast, error) {
	podcast := &entities.Podcast{
		PageLink: link,
	}

	if err := m.crawler.Find(podcast); err != nil {
		return nil, err
	}
	return podcast, m.downloader.Download(podcast)
}

// Start a full or partial flow of steps the bot is cable of
// it may just find and download the file and return the path or go down till translation
func (m Manager) FullFlow(link string) (*entities.Podcast, error) {
	podcast := &entities.Podcast{
		PageLink: link,
	}

	if err := m.crawler.Find(podcast); err != nil {
		return nil, err
	}

	if err := m.downloader.Download(podcast); err != nil {
		return nil, err
	}

	if err := m.convertor.Convert(podcast); err != nil {
		return nil, err
	}

	if err := m.transcriber.Transcribe(podcast); err != nil {
		return nil, err
	}

	return podcast, m.translator.Translate(podcast)
}

func (m Manager) FullExceptTranslation(link string) (*entities.Podcast, error) {
	podcast := &entities.Podcast{
		PageLink: link,
	}

	if err := m.crawler.Find(podcast); err != nil {
		return nil, err
	}

	if err := m.downloader.Download(podcast); err != nil {
		return nil, err
	}

	if err := m.convertor.Convert(podcast); err != nil {
		return nil, err
	}

	return podcast, m.transcriber.Transcribe(podcast)
}

func (m Manager) TranscribeDownloadedMP3(path string) (*entities.Podcast, error) {
	podcast := &entities.Podcast{
		Mp3Path: path,
	}

	if err := m.convertor.Convert(podcast); err != nil {
		return nil, err
	}

	return podcast, m.transcriber.Transcribe(podcast)
}

func (m Manager) TranslateDownloadedMP3(path string) (*entities.Podcast, error) {
	podcast := &entities.Podcast{
		Mp3Path: path,
	}

	if err := m.convertor.Convert(podcast); err != nil {
		return nil, err
	}

	if err := m.transcriber.Transcribe(podcast); err != nil {
		return nil, err
	}

	return podcast, m.translator.Translate(podcast)
}

func (m Manager) TranscribeDownloadedWAV(path string) (*entities.Podcast, error) {
	podcast := &entities.Podcast{
		WavPath: path,
	}

	return podcast, m.transcriber.Transcribe(podcast)
}

func (m Manager) TranslateDownloadedWAV(path string) (*entities.Podcast, error) {
	podcast := &entities.Podcast{
		WavPath: path,
	}

	if err := m.transcriber.Transcribe(podcast); err != nil {
		return nil, err
	}

	return podcast, m.translator.Translate(podcast)
}
