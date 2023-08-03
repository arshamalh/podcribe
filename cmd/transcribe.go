package cmd

import (
	"fmt"
	"path"
	"podcribe/manager"
	"podcribe/services/convertor"
	"podcribe/services/crawler"
	"podcribe/services/downloader"
	"podcribe/services/translator"
	"podcribe/services/whisper"
	"podcribe/tools"
	"strings"

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
		manager.FullFlow,
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
		ext := path.Ext(link)
		if ext == ".wav" {
			manager.TranslateDownloadedWAV(link)
		} else if ext == ".mp3" {
			manager.TranslateDownloadedMP3(link)
		}
	}
}
