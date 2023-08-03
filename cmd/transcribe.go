package cmd

import (
	"fmt"
	"os"
	"path"
	"podcribe/manager"
	"podcribe/services/convertor"
	"podcribe/services/crawler"
	"podcribe/services/downloader"
	"podcribe/services/translator"
	"podcribe/services/whisper"
	"podcribe/tools"
	"strings"
	"time"

	"github.com/spf13/cobra"
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
	manager := manager.New(
		crawler.New(),
		downloader.New(),
		convertor.New(),
		whisper.New(),
		translator.New(),
	)

	isOnWeb := strings.HasPrefix(link, "http://") || strings.HasPrefix(link, "https://")
	if isOnWeb {
		fmt.Println("started downloading:", link)
		translation, err := manager.FullFlow(link)

		if err != nil {
			fmt.Println(err)
		}

		filepath, err := tools.WriteTranslation("tempfilename", translation)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Podcast translation successfully generated in: ", filepath)
	} else {
		info, err := os.Stat(link) // TODO: Can we use file info to improve UX?
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
			manager.TranslateDownloadedWAV(link)
		} else if ext == ".mp3" {
			manager.TranslateDownloadedMP3(link)
		} else {
			fmt.Println("file format is not supported")
		}
	}
}
