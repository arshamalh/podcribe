package cmd

import (
	"fmt"
	"os"
	"path"
	"podcribe/log"
	"podcribe/manager"
	"podcribe/repo/sqlite"
	"podcribe/services/convertor"
	"podcribe/services/crawler"
	"podcribe/services/downloader"
	"podcribe/services/transcriber"
	"podcribe/services/translator"
	"podcribe/tools"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func registerTranscribeCmd(root *cobra.Command) {
	cmd := &cobra.Command{
		Use:     "transcribe",
		Short:   "transcribes (extract text out of voice) podcasts",
		Aliases: []string{"v2t"},
		Run: func(cmd *cobra.Command, args []string) {
			transcribe(args[0])
		},
	}

	root.AddCommand(cmd)
}

func transcribe(link string) {
	db, err := sqlite.New()
	if err != nil {
		fmt.Println(err)
	}

	manager := manager.New(
		crawler.New(db),
		downloader.New(db, 3),
		convertor.New(),
		transcriber.New(),
		translator.New(),
	)

	isOnWeb := strings.HasPrefix(link, "http://") || strings.HasPrefix(link, "https://")
	if isOnWeb {
		fmt.Println("started downloading:", link)
		podcast, err := manager.JustDownload(link)

		if err != nil {
			fmt.Println(err)
		}

		filepath, err := tools.WriteTranslation("tempfilename", podcast.TranscriptionPath)
		if err != nil {
			fmt.Println(err)
		}
		// TODO: make it a more clear message
		fmt.Printf("Podcast translation successfully generated in: %s, %#v", filepath, podcast)
	} else {
		info, err := os.Stat(link)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Println("file does not exist")
				return
			} else {
				fmt.Println("unexpected error in detecting file location")
				return
			}
		}
		ext := path.Ext(link)
		fmt.Printf("File detected!\nName: %s\nSize: %d Bytes\nLast Modified: %s\n*****\n\n", info.Name(), info.Size(), info.ModTime().Format(time.RFC822))
		if ext == ".wav" {
			manager.TranscribeDownloadedWAV(link)
		} else if ext == ".mp3" {
			_, err := manager.TranscribeDownloadedMP3(link)
			if err != nil {
				log.Info("can't transcribe", zap.Error(err))
			}
		} else {
			fmt.Println("file format is not supported")
		}
	}
}
