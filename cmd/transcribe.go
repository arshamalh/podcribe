package cmd

import (
	"podcribe/manager"
	"podcribe/services/convertor"
	"podcribe/services/crawler"
	"podcribe/services/downloader"
	"podcribe/services/translator"
	"podcribe/services/whisper"

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
	manager.Start(link)
}
