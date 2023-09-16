package repo

import "time"

type Podcast struct {
	Id                int
	PageLink          string
	Mp3Link           string
	Provider          string
	Mp3Path           string
	WavPath           string
	TranscriptionPath string
	TranslationPath   string
	ReferencedCount   int
	CreatedAt         time.Time
}

type User struct {
	Id        int
	CreatedAt time.Time
}

type UserPodcast struct {
	Id        int
	UserId    int
	PodcastId int
}

type DB interface {
	StorePodcast(podcast Podcast) (err error)
	GetPodcastByPageLink(pageLink string) (podcast Podcast, err error)
	IncreasePodcastReferencedCount(podcastId int) (err error)
}
