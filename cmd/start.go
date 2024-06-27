package cmd

import (
	"podcribe/config"
	"podcribe/log"
	"sync"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var (
	token       string
	openAIToken string
	openAIBase  string
	server      bool
	telegramOn  bool
	port        int
)

func registerStart(root *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "starting telegram bot",
		Run: func(cmd *cobra.Command, args []string) {
			log.Initialize()
			cfg := *config.Setter().
				SetTelegramToken(token).
				SetOpenAIToken(openAIToken).
				SetOpenAIBase(openAIBase)
			start(cfg)
		},
	}

	cmd.Flags().BoolVarP(&server, "server", "s", false, "whether we should start a webserver or not")
	cmd.Flags().IntVarP(&port, "port", "p", 8080, "port for api server")

	cmd.Flags().StringVarP(&token, "telegram-token", "t", "", "input your telegram token") // TODO: Start telegram or webserver or both, in case of neither, throw an error
	cmd.Flags().BoolVar(&telegramOn, "telegram-on", false, "whether the telegram should be on or not")

	cmd.Flags().StringVar(&openAIToken, "open-ai-token", "", "token for connecting to OpenAI")
	cmd.Flags().StringVar(&openAIBase, "open-ai-base", "", "address of OpenAI or its proxy")
	root.AddCommand(cmd)
}

func start(cfg config.Config) {
	if err := godotenv.Load(); err != nil {
		log.Gl.Error(err.Error())
	}

	if telegramOn && cfg.TelegramToken == "" {
		log.Gl.Fatal("no telegram token provided, no server setting provided, there is nothing to start")
	}

	var wg sync.WaitGroup
	if telegramOn {
		wg.Add(1)
		go startTelegram(cfg, &wg)
	}
	if server {
		wg.Add(1)
		go startAPI(port, &wg)
	}
	wg.Wait()
}
