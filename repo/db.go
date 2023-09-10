package repo

import "time"

type Podcast struct {
	Id              int
	PageLink        string
	PodcastLink     string
	Provider        string
	Path            string
	ReferencedCount int
	CreatedAt       time.Time
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
	StorePodcast(Podcast) error
}
