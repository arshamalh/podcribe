package downloader_test

import (
	"fmt"
	"podcribe/entities"
	"podcribe/services/downloader"
	"testing"
)

func TestDownloader(t *testing.T) {
	downloader := downloader.New(3)
	podcast := new(entities.Podcast)
	podcast.Mp3Link = "https://dts.podtrac.com/redirect.mp3/chrt.fm/track/18987/api.spreaker.com/download/episode/56636674/2070_0905.mp3"
	err := downloader.Download(podcast)
	// filepath, err := downloader.Download("https://dts.podtrac.com/redirect.mp3/chrt.fm/track/18987/api.spreaker.com/download/episode/56636005/ep2070.mp3")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(podcast.Mp3Path)
}

// For future test case
// jsdata="Kwyn5e;https://dts.podtrac.com/redirect.mp3/chrt.fm/track/18987/api.spreaker.com/download/episode/56486403/0821_extra.mp3;$56"

// Head "https://api.spreaker.com/download/episode/56486403/0821_extra.mp3": read tcp 192.168.80.121:55404->10.10.34.35:443: read: connection reset by peer

// Another weired link:
// https://dts.podtrac.com/redirect.mp3/download.softskills.audio/sse-369.mp3
// but downloadable
