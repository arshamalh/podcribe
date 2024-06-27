package manager

import (
	"path"
	"podcribe/entities"
	"podcribe/services/convertor"
	"podcribe/services/crawler"
	"podcribe/services/downloader"
	"podcribe/services/transcriber"
	"podcribe/services/translator"
	"strings"
)

// We may want to just download the podcast using the bot (no translating, convertion)

// Holds different services which each has its own settings
// Make a new manager for each user
// Can hold its own settings too
type Manager struct {
	crawler     *crawler.CrawlHandler
	downloader  downloader.I
	convertor   convertor.I
	transcriber transcriber.I
	translator  translator.I
}

func New(crawler *crawler.CrawlHandler, downloader downloader.I, convertor convertor.I, transcriber transcriber.I, translator translator.I) *Manager {
	return &Manager{
		crawler, downloader, convertor, transcriber, translator,
	}
}

// the only difference is filtered providers
// TODO: Should we include a direct download for this function?
func (m Manager) JustDownload(link string) (*entities.Podcast, error) {
	mp3Link, err := m.crawler.Find(link)
	if err != nil {
		return nil, err
	}

	podcast := &entities.Podcast{
		PageLink: link,
		Mp3Link:  mp3Link,
	}

	return podcast, m.downloader.Download(podcast)
}

// Start a full or partial flow of steps the bot is cable of
// it may just find and download the file and return the filepath or go down till translation
func (m Manager) FullFlow(link string) (*entities.Podcast, error) {
	mp3Link, err := m.crawler.Find(link)
	if err != nil {
		return nil, err
	}

	podcast := &entities.Podcast{
		PageLink: link,
		Mp3Link:  mp3Link,
	}

	if err := m.downloader.Download(podcast); err != nil {
		return nil, err
	}

	if err := m.convertor.ConvertMP3ToWAV(podcast); err != nil {
		return nil, err
	}

	if err := m.transcriber.Transcribe(podcast); err != nil {
		return nil, err
	}

	return podcast, m.translator.Translate(podcast)
}

func (m Manager) FullExceptTranslation(link string) (*entities.Podcast, error) {
	mp3Link, err := m.crawler.Find(link)
	if err != nil {
		return nil, err
	}

	podcast := &entities.Podcast{
		PageLink: link,
		Mp3Link:  mp3Link,
	}

	if err := m.downloader.Download(podcast); err != nil {
		return nil, err
	}

	if err := m.convertor.ConvertMP3ToWAV(podcast); err != nil {
		return nil, err
	}

	return podcast, m.transcriber.Transcribe(podcast)
}

func (m Manager) TranscribeDownloadedMP3(filepath string) (*entities.Podcast, error) {
	_, filename := path.Split(filepath)
	filename, _ = strings.CutSuffix(filename, ".mp3")
	podcast := &entities.Podcast{
		Name:    filename,
		Mp3Path: filepath,
	}

	if err := m.convertor.ConvertMP3ToWAV(podcast); err != nil {
		return nil, err
	}

	return podcast, m.transcriber.Transcribe(podcast)
}

func (m Manager) TranslateDownloadedMP3(filepath string) (*entities.Podcast, error) {
	podcast := &entities.Podcast{
		Mp3Path: filepath,
	}

	if err := m.convertor.ConvertMP3ToWAV(podcast); err != nil {
		return nil, err
	}

	if err := m.transcriber.Transcribe(podcast); err != nil {
		return nil, err
	}

	return podcast, m.translator.Translate(podcast)
}

func (m Manager) TranscribeDownloadedWAV(filepath string) (*entities.Podcast, error) {
	podcast := &entities.Podcast{
		WavPath: filepath,
	}

	return podcast, m.transcriber.Transcribe(podcast)
}

func (m Manager) TranslateDownloadedWAV(filepath string) (*entities.Podcast, error) {
	podcast := &entities.Podcast{
		WavPath: filepath,
	}

	if err := m.transcriber.Transcribe(podcast); err != nil {
		return nil, err
	}

	return podcast, m.translator.Translate(podcast)
}
