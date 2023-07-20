package crawler

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
	return "", nil
}
