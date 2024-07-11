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
	InvoiceFactor float64
	Wallets       struct {
		TRON       string
		TON        string
		USDT_TRC20 string
	}

	TRON_APIKey string
	TON_APIKey  string
	TRON_BASE   string
	TON_BASE    string

	WallexBase string
	WallexKey  string
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

	config.TRON_APIKey = os.Getenv("TRON_API_KEY")
	config.TON_APIKey = os.Getenv("TON_API_KEY")

	config.WallexKey = os.Getenv("WALLEX_KEY")
}

func setDefaultValues() {
	config.OpenAIBase = "https://api.gilas.io/v1"
	config.AdminUsername = "@arshamalh"
	config.InvoiceFactor = 0.001 // 1/1000 $ per second
	config.Wallets.TON = "UQBk1vUdJyOuYw0jY8hPILuJdjbJ2BLAqpmhA_IzewQX0tKb"
	config.Wallets.TRON = "TKQpSBWAQF1434J2gCoAN5WPALoft9yMhc"
	config.Wallets.USDT_TRC20 = "TKQpSBWAQF1434J2gCoAN5WPALoft9yMhc"

	config.TON_BASE = "https://tonapi.io"
	config.TRON_BASE = "https://apilist.tronscanapi.com"

	config.WallexBase = "https://api.wallex.ir"
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
