package crawler

import (
	"strings"

	"github.com/gocolly/colly/v2"
)

type googleCrawler struct {
}

func NewGoogleCrawler() *googleCrawler {
	return &googleCrawler{}
}

// Get a page link as an input and search in the page for a podcast link,
// returns podcast link if any or raise an error if there is no link or the page is not accessible.
func (gc googleCrawler) Find(pageLink string) (string, error) {
	cl := colly.NewCollector()
	podcastLinks := make([]string, 0)

	cl.OnHTML(`div[jsname="fvi9Ef"][jsdata]`, func(e *colly.HTMLElement) {
		jsdata := e.Attr("jsdata")
		httpIndex := strings.LastIndex(jsdata, "https")
		mp3Index := strings.Index(jsdata, ".mp3")
		if mp3Index > httpIndex {
			podcastLinks = append(podcastLinks, jsdata[httpIndex:mp3Index+4])
		}
	})

	if err := cl.Visit(pageLink); err != nil {
		return "", err
	}

	return podcastLinks[0], nil
}

// page_link := podcast.PageLink
// 	if strings.Contains(page_link, "google") {
// 		podcastLinks, err := getGooglePodcastLinks(page_link)
// 		if err != nil {
// 			return err
// 		}
// 		podcast.Mp3Link = podcastLinks[0]
// 		return nil
// 	}

// 	if strings.Contains(page_link, "castbox") {
// 		podcastLinks, err := getCastboxPodcastLinks(page_link)
// 		if err != nil {
// 			return err
// 		}
// 		podcast.Mp3Link = podcastLinks[0]
// 		return nil
// 	}

// 	return errors.New("unknown provider")
