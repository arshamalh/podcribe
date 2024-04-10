package cmd

import (
	"os"
	"podcribe/log"
	"sync"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var (
	token      string
	server     bool
	telegramOn bool
	port       int
)

func registerStart(root *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "starting telegram bot",
		Run: func(cmd *cobra.Command, args []string) {
			log.Initialize()
			start()
		},
	}

	cmd.Flags().BoolVarP(&server, "server", "s", false, "this flag shows whether we should start a webserver or not")
	cmd.Flags().StringVarP(&token, "token", "t", "", "input your telegram token") // TODO: Start telegram or webserver or both, in case of neither, throw an error
	cmd.Flags().BoolVar(&telegramOn, "telegram-on", false, "whether the telegram should be on or not")
	cmd.Flags().IntVarP(&port, "port", "p", 8080, "port for api server")
	root.AddCommand(cmd)
}

func start() {
	if err := godotenv.Load(); err != nil {
		log.Gl.Error(err.Error())
	}

	if token == "" {
		if os.Getenv("TOKEN") != "" {
			token = os.Getenv("TOKEN")
		} else if server {
			// TODO: Start server is empty and redundant,
			// but startAPI also seems to be in the wrong place,
			// both of them should run in separate go-routines,
			// And synchronized using wait groups
		} else {
			log.Gl.Fatal("no telegram token provided, no server setting provided, there is nothing to start")
		}
	}
	var wg sync.WaitGroup
	if telegramOn {
		wg.Add(1)
		go startTelegram(token, &wg)
	}
	if server {
		startAPI(port)
	}
	wg.Wait()
}
