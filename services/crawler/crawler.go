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
	// Check if the podcast already exists in the database
	podcastModel, exist, err := c.isPodcastExist(podcast.PageLink)
	if err != nil {
		return err
	}
	if exist {
		// Increase the referenced count and set the mp3 link
		err = c.db.IncreasePodcastReferencedCount(podcastModel.Id)
		if err != nil {
			return err
		}
		podcast.Id = podcastModel.Id
		podcast.Mp3Link = podcastModel.Mp3Link
		return nil
	}

	// Get the podcast mp3 link based on the provider
	mp3Link, provider, err := getProviderMp3PodcastLink(podcast.PageLink)
	if err != nil {
		return err
	}

	// Store the new podcast in the database
	newPodcastModel := podcast.Model(map[string]any{
		"provider": provider,
		"mp3_link": mp3Link,
	})
	err = c.db.StorePodcast(newPodcastModel)
	if err != nil {
		return err
	}

	podcast.Mp3Link = mp3Link

	return nil
}

func (c Crawler) isPodcastExist(PageLink string) (podcastModel repo.Podcast, exist bool, err error) {
	podcastModel, err = c.db.GetPodcastByPageLink(PageLink)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return podcastModel, false, nil
		}
		return podcastModel, false, err
	}

	return podcastModel, true, nil
}

func getProviderMp3PodcastLink(pageLink string) (mp3Link, provider string, err error) {

	if strings.Contains(pageLink, GOOGLE) {
		return getGoogleMp3Links(pageLink)
	}

	if strings.Contains(pageLink, CASTBOX) {
		return getCastboxMp3Links(pageLink)
	}

	return "", "", errors.New("unknown provider")
}

// Find google podcast .mp3 link
func getGoogleMp3Links(pageLink string) (podcastMp3Link, provider string, err error) {
	cl := colly.NewCollector()

	var podcastMp3Links []string
	cl.OnHTML(`div[jsname="fvi9Ef"][jsdata]`, func(e *colly.HTMLElement) {
		jsdata := e.Attr("jsdata")
		httpIndex := strings.LastIndex(jsdata, "https")
		mp3Index := strings.Index(jsdata, ".mp3")
		if mp3Index > httpIndex {
			podcastMp3Links = append(podcastMp3Links, jsdata[httpIndex:mp3Index+4])
		}
	})

	err = cl.Visit(pageLink)
	if err != nil {
		return "", "", err
	}

	return podcastMp3Links[0], GOOGLE, nil
}

// Find castbox .mp3 link
func getCastboxMp3Links(pageLink string) (podcastMp3Link, provider string, err error) {
	cl := colly.NewCollector()

	var podcastMp3Links []string
	cl.OnHTML(`#root > div > div:nth-child(1) > audio > source `, func(e *colly.HTMLElement) {
		link := e.Attr("src")
		httpIndex := strings.LastIndex(link, "https")
		mp3Index := strings.Index(link, ".mp3")
		if mp3Index > httpIndex {
			podcastMp3Links = append(podcastMp3Links, link[httpIndex:mp3Index+4])
		}
	})

	err = cl.Visit(pageLink)
	if err != nil {
		return "", "", err
	}

	return podcastMp3Links[0], CASTBOX, nil
}
