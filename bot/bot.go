package bot

import (
	"log"

	"github.com/arian-press2015/apcore_bot/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBot struct {
	Bot    *tgbotapi.BotAPI
	Config *config.Config
}

func NewTelegramBot(config *config.Config) (*TelegramBot, error) {
	botToken := config.BotToken
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return nil, err
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	webhook, err := tgbotapi.NewWebhook(config.WebhookURL + bot.Token)
	if err != nil {
		log.Panic(err)
	}

	_, err = bot.Request(webhook)
	if err != nil {
		log.Panic(err)
	}

	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}

	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}

	return &TelegramBot{Bot: bot, Config: config}, nil
}

func (b *TelegramBot) GetWebhookURL() string {
	return "/" + b.Bot.Token
}
