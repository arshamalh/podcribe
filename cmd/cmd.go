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

	// http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	registerStart(root)
	registerTranscribe(root)
	registerMigrate(root)

	if err := root.Execute(); err != nil {
		fmt.Println(err)
	}
}
