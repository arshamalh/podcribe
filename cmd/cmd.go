package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Execute() {

	var root = &cobra.Command{
		Use:   "podcribe",
		Short: "crawl, download, transcribe and translates podcasts",
	}

	registerStartCmd(root)
	registerTranscribeCmd(root)

	if err := root.Execute(); err != nil {
		fmt.Println(err)
	}
}
