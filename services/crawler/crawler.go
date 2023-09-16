package crawler

import (
	"database/sql"
	"errors"
	"github.com/gocolly/colly/v2"
	"podcribe/entities"
	"podcribe/repo"
	"strings"
)

const GOOGLE = "google"
const CASTBOX = "castbox"

type I interface {
	// Get a page link as an input and search in the page for a podcast link,
	// returns podcast link if any or raise an error if there is no link or the page is not accessible.
	Find(podcast *entities.Podcast) error
}

type Crawler struct {
	db repo.DB
}

func New(db repo.DB) *Crawler {
	return &Crawler{db}
}

// Get a page link as an input and search in the page for a podcast link,
// returns podcast link if any or raise an error if there is no link or the page is not accessible.
func (c Crawler) Find(podcast *entities.Podcast) error {
	podcastModel1, err := c.db.GetPodcastByPageLink(podcast.PageLink)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			mp3Link, provider, err := getPodcastLink(podcast.PageLink)
			if err != nil {
				return err
			}

			podcastModel := podcast.Model(map[string]any{
				"provider": provider,
				"mp3_link": mp3Link,
			})
			err = c.db.StorePodcast(podcastModel)
			if err != nil {
				return err
			}

			podcast.Mp3Link = mp3Link

			return nil
		}
		return err
	}

	err = c.db.IncreasePodcastReferencedCount(podcastModel1.Id)
	if err != nil {
		return err
	}
	podcast.Mp3Link = podcastModel1.Mp3Link

	return nil
}

// Find podcast link
func getPodcastLink(page_link string) (podcastLink string, provider string, err error) {
	var podcastLinks []string
	if strings.Contains(page_link, GOOGLE) {
		podcastLinks, err = getGooglePodcastLinks(page_link)
		if err != nil {
			return "", "", err
		}
		return podcastLinks[0], GOOGLE, nil
	}

	if strings.Contains(page_link, CASTBOX) {
		podcastLinks, err = getCastboxPodcastLinks(page_link)
		if err != nil {
			return "", "", err
		}
		return podcastLinks[0], CASTBOX, nil
	}

	return "", "", errors.New("unknown provider")
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
