package cmd

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

func Execute() {
	var root = &cobra.Command{
		Use:   "podcribe",
		Short: "crawl, download, transcribe and translates podcasts",
	}

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	registerStartCmd(root)
	registerTranscribeCmd(root)

	if err := root.Execute(); err != nil {
		fmt.Println(err)
	}
}
