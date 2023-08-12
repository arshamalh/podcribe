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
	"strings"
	"sync"
)

// TODO: Show the downloading progress in Telegram, UI, CLI or whatever is using this application

type I interface {
	// Downloads a file, store it in the file system and returns the path to the file,
	// or raise an error if it can't download the file or can't store it.
	Download(url string) (filePath string, err error)
}

type Downloader struct {
	workersCount int
	chunks       map[int][]byte
	mx           sync.Mutex
}

func New() *Downloader {
	return &Downloader{
		workersCount: 3,
		chunks:       make(map[int][]byte),
	}
}

func (d Downloader) Download(uri string) (filePath string, err error) {
	fmt.Println("uri	", uri)
	isSupported, contentLength, err := getRangeDetails(uri)
	if err != nil {
		return "", err
	}

	if !isSupported || d.workersCount <= 1 {
		return d.processSingle(uri)
	}

	return d.processMultiple(contentLength, uri)
}

func (d Downloader) processSingle(uri string) (filePath string, err error) {
	//Initialize first index with []byte
	d.chunks[0] = make([]byte, 0)
	d.downloadFileForRange(nil, uri, "", 0)

	if err != nil {
		return "", err
	}

	return d.combineChunks(uri)
}

func (d Downloader) processMultiple(contentLength int, uri string) (filePath string, err error) {
	split := contentLength / d.workersCount
	var wg sync.WaitGroup
	index := 0

	for i := 0; i < contentLength; i += split + 1 {
		j := i + split
		if j > contentLength {
			j = contentLength
		}

		d.chunks[index] = make([]byte, 0)
		wg.Add(1)
		go d.downloadFileForRange(&wg, uri, strconv.Itoa(i)+"-"+strconv.Itoa(j), index)
		index++
	}

	wg.Wait()

	if err != nil {
		return "", err
	}

	return d.combineChunks(uri)
}

func (d *Downloader) downloadFileForRange(wg *sync.WaitGroup, uri, r string, index int) {
	defer wg.Done()

	request, err := http.NewRequest("GET", uri, strings.NewReader(""))
	if err != nil {
		return
	}

	if r != "" {
		request.Header.Add("Range", "bytes="+r)
	}

	sc, _, data, err := doAPICall(request)
	if err != nil {
		return
	}

	if sc != 200 && sc != 206 {
		err = fmt.Errorf("invalid status code: %d", sc)
		return
	}

	d.mx.Lock()
	d.chunks[index] = append(d.chunks[index], data...)
	d.mx.Unlock()
}

func (d Downloader) combineChunks(uri string) (filePath string, err error) {
	currDir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	filePath = path.Join(currDir, "/", filepath.Base(uri))

	out, err := os.Create(filePath)
	defer out.Close()

	if err != nil {
		return "", err
	}

	buf := bytes.NewBuffer(nil)
	for i := 0; i < len(d.chunks); i++ {
		buf.Write(d.chunks[i])
	}

	l, err := buf.WriteTo(out)
	if err != nil {
		return "", err
	}

	fmt.Println("Wrote to File : %v, len : %v", filePath, l)

	return filePath, nil
}

func getRangeDetails(uri string) (bool, int, error) {
	request, err := http.NewRequest("HEAD", uri, strings.NewReader(""))
	if err != nil {
		return false, 0, err
	}

	sc, headers, _, err := doAPICall(request)
	if err != nil {
		return false, 0, err
	}

	if sc != 200 && sc != 206 {
		return false, 0, err
	}

	conLen := headers.Get("Content-Length")
	cl, err := strconv.Atoi(conLen)
	if err != nil {
		return false, 0, err
	}

	if headers.Get("Accept-Ranges") == "bytes" {
		return true, cl, nil
	}

	return false, cl, nil
}

func doAPICall(request *http.Request) (statusCode int, header http.Header, data []byte, err error) {
	client := http.Client{
		Timeout: 0,
	}

	response, err := client.Do(request)
	if err != nil {
		return 0, http.Header{}, []byte{}, err
	}
	defer response.Body.Close()

	data, err = io.ReadAll(response.Body)
	if err != nil {
		return 0, http.Header{}, []byte{}, err
	}

	return response.StatusCode, response.Header, data, nil
}
