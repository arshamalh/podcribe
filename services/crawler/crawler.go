package crawler

import (
	"errors"
	"strings"
)

type Provider interface {
	Find(pageLink string) (string, error)
}

type CrawlHandler struct {
	providers map[string]Provider
}

func New() *CrawlHandler {
	return &CrawlHandler{
		providers: make(map[string]Provider),
	}
}

func (ch CrawlHandler) RegisterCrawler(providerName string, provider Provider) {
	ch.providers[providerName] = provider
}

func (ch CrawlHandler) RegisterDefaults() {
	ch.providers["google"] = NewGoogleCrawler()
	ch.providers["castbox"] = NewCastboxCrawler()
}

// Get a page link as an input and search in the page for a podcast link,
// returns podcast link if any or raise an error if there is no link or the page is not accessible.
func (ch CrawlHandler) Find(pageLink string) (string, error) {
	for providerName, provider := range ch.providers {
		if strings.Contains(pageLink, providerName) {
			mp3Link, err := provider.Find(pageLink)
			if err != nil {
				return "", err
			}
			return mp3Link, nil
		}
	}

	return "", errors.New("unknown provider")
}
