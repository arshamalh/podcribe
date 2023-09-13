package crawler

import (
	"errors"
	"podcribe/entities"
	"strings"

	"github.com/gocolly/colly/v2"
)

type I interface {
	// Get a page link as an input and search in the page for a podcast link,
	// returns podcast link if any or raise an error if there is no link or the page is not accessible.
	Find(podcast *entities.Podcast) error
}

type crawler struct {
}

func New() *crawler {
	return &crawler{}
}

// Get a page link as an input and search in the page for a podcast link,
// returns podcast link if any or raise an error if there is no link or the page is not accessible.
func (c crawler) Find(podcast *entities.Podcast) error {
	page_link := podcast.PageLink
	if strings.Contains(page_link, "google") {
		podcastLinks, err := getGooglePodcastLinks(page_link)
		if err != nil {
			return err
		}
		podcast.Mp3Link = podcastLinks[0]
		return nil
	}

	if strings.Contains(page_link, "castbox") {
		podcastLinks, err := getCastboxPodcastLinks(page_link)
		if err != nil {
			return err
		}
		podcast.Mp3Link = podcastLinks[0]
		return nil
	}

	return errors.New("unknown provider")
}

// Find google podcast .mp3 link
func getGooglePodcastLinks(page_link string) (podcast_link []string, err error) {
	cl := colly.NewCollector()

	cl.OnHTML(`div[jsname="fvi9Ef"][jsdata]`, func(e *colly.HTMLElement) {
		jsdata := e.Attr("jsdata")
		httpIndex := strings.LastIndex(jsdata, "https")
		mp3Index := strings.Index(jsdata, ".mp3")
		if mp3Index > httpIndex {
			podcast_link = append(podcast_link, jsdata[httpIndex:mp3Index+4])
		}
	})

	err = cl.Visit(page_link)
	if err != nil {
		return nil, err
	}

	return podcast_link, nil
}

// Find castbox .mp3 link
func getCastboxPodcastLinks(page_link string) (podcast_link []string, err error) {
	cl := colly.NewCollector()

	cl.OnHTML(`#root > div > div:nth-child(1) > audio > source `, func(e *colly.HTMLElement) {
		link := e.Attr("src")
		httpIndex := strings.LastIndex(link, "https")
		mp3Index := strings.Index(link, ".mp3")
		if mp3Index > httpIndex {
			podcast_link = append(podcast_link, link[httpIndex:mp3Index+4])
		}
	})

	err = cl.Visit(page_link)
	if err != nil {
		return nil, err
	}

	return podcast_link, nil
}
