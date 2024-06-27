package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TelegramToken string
	OpenAIToken   string
	OpenAIBase    string
	AdminUsername string
}

var config = new(Config)

func init() {
	setDefaultValues()
	if err := godotenv.Load(); err != nil {
		log.Println(err.Error())
	}
	config.SetTelegramToken(os.Getenv("TELEGRAM_TOKEN"))
	config.SetOpenAIBase(os.Getenv("OPENAI_BASE"))
	config.SetOpenAIToken(os.Getenv("OPENAI_TOKEN"))
}

func setDefaultValues() {
	config.OpenAIBase = "https://api.gilas.io/v1"
	config.AdminUsername = "@arshamalh"
}

func Get() Config {
	return *config
}

func Setter() *Config {
	return config
}

func (cfg *Config) SetTelegramToken(token string) *Config {
	if token != "" {
		cfg.TelegramToken = token
	}
	return cfg
}

func (cfg *Config) SetOpenAIBase(base string) *Config {
	if base != "" {
		cfg.OpenAIBase = base
	}
	return cfg
}

func (cfg *Config) SetOpenAIToken(token string) *Config {
	if token != "" {
		cfg.OpenAIToken = token
	}
	return cfg
}
