package downloader

// TODO: Show the downloading progress in Telegram, UI, CLI or whatever is using this application

type I interface {
	// Downloads a file, store it in the file system and returns the path to the file,
	// or raise an error if it can't download the file or can't store it.
	Download(url string) (filepath string, err error)
}

type downloader struct {
}

func New() *downloader {
	return &downloader{}
}

func (d downloader) Download(url string) (filepath string, err error) {
	return "", nil
}
