package downloader

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"sync"
)

// TODO: Show the downloading progress in Telegram, UI, CLI or whatever is using this application

type I interface {
	// Downloads a file, store it in the file system and returns the path to the file,
	// or raise an error if it can't download the file or can't store it.
	Download(url string) (filePath string, err error)
}

type downloader struct {
	workersCount int // TODO Calculate workers count dynamically and combine its logic with process single
	chunks       [][]byte
}

func New(workers_count int) *downloader {
	return &downloader{
		workersCount: workers_count,
		chunks:       make([][]byte, workers_count),
	}
}

func (d *downloader) Download(uri string) (filePath string, err error) {
	fmt.Println("uri	", uri)
	isSupported, contentLength, err := getRangeDetails(uri)
	if err != nil {
		return "", err
	}

	if !isSupported || d.workersCount <= 1 {
		return d.processSingle(uri)
	}

	filePath, err = d.processMultiple(contentLength, uri)
	if err != nil {
		return "", nil
	}

	fmt.Printf("Wrote to File : %v, len : %v\n", filePath, contentLength)
	return filePath, nil
}

func (d *downloader) processSingle(uri string) (filePath string, err error) {
	//Initialize first index with []byte
	d.chunks[0] = make([]byte, 0)
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

	request, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return
	}

	if _range != "" {
		request.Header.Add("Range", "bytes="+_range)
	}

	sc, _, data, err := doAPICall(request)
	if err != nil {
		return
	}

	if sc != 200 && sc != 206 {
		// err = fmt.Errorf("invalid status code: %d", sc)
		return
	}

	d.chunks[index] = make([]byte, 0)
	d.chunks[index] = append(d.chunks[index], data...)
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

	buf := bytes.NewBuffer(nil)
	for i := 0; i < len(d.chunks); i++ {
		if _, err := buf.Write(d.chunks[i]); err != nil {
			return "", err
		}
	}

	if _, err = buf.WriteTo(output); err != nil {
		return "", err
	}

	return filePath, nil
}

func getRangeDetails(uri string) (bool, int, error) {
	request, err := http.NewRequest("HEAD", uri, nil)
	if err != nil {
		return false, 0, err
	}

	sc, headers, _, err := doAPICall(request)
	if err != nil {
		// If resets by peer, we should tell user that we don't support downloading this podcast
		return false, 0, err
	}

	if sc != 200 && sc != 206 {
		return false, 0, err
	}

	contentLength, err := strconv.Atoi(headers.Get("Content-Length"))
	if err != nil {
		return false, 0, err
	}

	if headers.Get("Accept-Ranges") == "bytes" {
		return true, contentLength, nil
	}

	return false, contentLength, nil
}

func doAPICall(request *http.Request) (statusCode int, header http.Header, data []byte, err error) {
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return 0, nil, nil, err
	}
	defer response.Body.Close()

	data, err = io.ReadAll(response.Body)
	if err != nil {
		return 0, nil, nil, err
	}

	return response.StatusCode, response.Header, data, nil
}
