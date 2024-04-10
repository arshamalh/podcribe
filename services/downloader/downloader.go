package downloader

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"podcribe/entities"
	"podcribe/log"
	"strconv"
	"sync"
	"time"

	"go.uber.org/zap"
)

type I interface {
	// Downloads a file, store it in the file system and returns the path to the file,
	// or raise an error if it can't download the file or can't store it.
	Download(podcast *entities.Podcast) error
}

type downloader struct {
	workersCount int // TODO Calculate workers count dynamically and combine its logic with process single
	chunks       []bytes.Buffer
	progressChan chan int
	// IMPORTANT
	// If multiple users use the same manager
	// If each user have a manager but each user can download multiple files at the same time
	// This logic won't work!
	// We can also have one manager per request and that kind of make sense
}

func New(workers_count int) *downloader {
	return &downloader{
		workersCount: workers_count,
		chunks:       make([]bytes.Buffer, workers_count),
		progressChan: make(chan int),
	}
}

func (d *downloader) Download(podcast *entities.Podcast) error {
	log.Gl.Info("downloading podcast", zap.String("uri", podcast.Mp3Link))
	isSupported, contentLength, err := getRangeDetails(podcast.Mp3Link)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go d.progress(ctx, contentLength)

	if !isSupported || d.workersCount <= 1 {
		fmt.Println("processing single")
		podcast.Mp3Path, err = d.processSingle(podcast.Mp3Link)
	} else {
		podcast.Mp3Path, err = d.processMultiple(contentLength, podcast.Mp3Link)
	}

	fmt.Printf("Wrote to File : %v, len : %v\n", podcast.Mp3Path, contentLength)
	return err
}

func (d *downloader) ConsumeProgress() <-chan int {
	return d.progressChan
}

func (d *downloader) processSingle(uri string) (filePath string, err error) {
	d.chunks[0] = bytes.Buffer{}
	d.downloadFileForRange(nil, uri, "", 0)

	if err != nil {
		return "", err
	}

	return d.combineChunks(uri)
}

func (d *downloader) processMultiple(contentLength int, uri string) (filePath string, err error) {
	partLength := contentLength / d.workersCount
	var wg sync.WaitGroup
	wg.Add(d.workersCount)

	for startRange, index := 0, 0; startRange < contentLength; startRange += partLength + 1 {
		endRange := startRange + partLength
		if endRange > contentLength {
			endRange = contentLength
		}
		_range := fmt.Sprintf("%d-%d", startRange, endRange)
		go d.downloadFileForRange(&wg, uri, _range, index)
		index++
	}

	wg.Wait()

	if err != nil {
		return "", err
	}

	return d.combineChunks(uri)
}

func (d *downloader) downloadFileForRange(wg *sync.WaitGroup, uri, _range string, index int) {
	defer wg.Done()
	fmt.Printf("\nrange %s started", _range)
	request, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return
	}

	if _range != "" {
		request.Header.Add("Range", "bytes="+_range)
	}

	response, err := (&http.Client{}).Do(request)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	fmt.Println("started writing to buffer")
	d.chunks[index] = bytes.Buffer{}
	written, err := io.Copy(&d.chunks[index], response.Body)
	fmt.Println(written, err)
}

func (d *downloader) combineChunks(uri string) (filePath string, err error) {
	currDir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	filePath = path.Join(currDir, "/", filepath.Base(uri))

	output, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer output.Close()

	for i := 0; i < len(d.chunks); i++ {
		if _, err = d.chunks[i].WriteTo(output); err != nil {
			return "", err
		}
	}

	return filePath, nil
}

func (d *downloader) progress(ctx context.Context, totalLen int) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			totalDownloaded := 0
			for _, chunk := range d.chunks {
				totalDownloaded += int((float32(chunk.Len()) / float32(totalLen)) * 100)
			}
			if totalDownloaded > 100 {
				totalDownloaded = 100
			}
			d.progressChan <- totalDownloaded
		}
		time.Sleep(time.Millisecond * 500)
	}
}

func getRangeDetails(uri string) (bool, int, error) {
	response, err := (&http.Client{}).Head(uri)

	if err != nil {
		// If resets by peer, we should tell user that we don't support downloading this podcast
		return false, 0, err
	}

	if response.StatusCode != 200 && response.StatusCode != 206 {
		return false, 0, err
	}

	contentLength, err := strconv.Atoi(response.Header.Get("Content-Length"))
	if err != nil {
		return false, 0, err
	}

	if response.Header.Get("Accept-Ranges") == "bytes" {
		return true, contentLength, nil
	}

	return false, contentLength, nil
}
