package crawler

import (
	"github.com/gocolly/colly/v2"
	"strings"
)

type I interface {
	// Get a page link as an input and search in the page for a podcast link,
	// returns podcast link if any or raise an error if there is no link or the page is not accessible.
	Find(page_link string) (podcast_link string, err error)
}

type crawler struct {
}

func New() *crawler {
	return &crawler{}
}

// Get a page link as an input and search in the page for a podcast link,
// returns podcast link if any or raise an error if there is no link or the page is not accessible.
func (c crawler) Find(page_link string) (podcast_link string, err error) {

	podcastLinks, err := getGooglePodcastLinks(page_link)
	if err != nil {
		return "", err
	}

	return podcastLinks[0], nil
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
