package cmd

import (
	"time"

	"github.com/spf13/cobra"
)

func registerStartCmd(root *cobra.Command) {
	cmd := &cobra.Command{
		Use: "start",
		Run: func(cmd *cobra.Command, args []string) {
			start()
		},
	}

	root.AddCommand(cmd)

}

func start() {
	for {
		time.Sleep(time.Hour)
	}
}
