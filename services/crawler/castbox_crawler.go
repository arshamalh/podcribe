package crawler

import (
	"strings"

	"github.com/gocolly/colly/v2"
)

type castboxCrawler struct {
}

func NewCastboxCrawler() *castboxCrawler {
	return &castboxCrawler{}
}

// Get a page link as an input and search in the page for a podcast link,
// returns podcast link if any or raise an error if there is no link or the page is not accessible.
func (cc castboxCrawler) Find(pageLink string) (string, error) {
	cl := colly.NewCollector()

	podcastLinks := make([]string, 0)

	cl.OnHTML(`#root > div > div:nth-child(1) > audio > source `, func(e *colly.HTMLElement) {
		link := e.Attr("src")
		httpIndex := strings.LastIndex(link, "https")
		mp3Index := strings.Index(link, ".mp3")
		if mp3Index > httpIndex {
			podcastLinks = append(podcastLinks, link[httpIndex:mp3Index+4])
		}
	})

	if err := cl.Visit(pageLink); err != nil {
		return "", err
	}

	return podcastLinks[0], nil
}
