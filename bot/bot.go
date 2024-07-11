package bot

import (
	"fmt"
	"log"
	"net/http"

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
	return &TelegramBot{Bot: bot, Config: config}, nil
}

func (b *TelegramBot) Run(handler func(http.ResponseWriter, *http.Request)) {
	webhook, err := tgbotapi.NewWebhook(b.Config.WebhookURL + b.Bot.Token)
	if err != nil {
		log.Panic(err)
	}

	_, err = b.Bot.Request(webhook)
	if err != nil {
		log.Panic(err)
	}

	info, err := b.Bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}

	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}

	updates := b.Bot.ListenForWebhook("/" + b.Bot.Token)
	http.HandleFunc("/"+b.Bot.Token, handler)
	go http.ListenAndServe(fmt.Sprintf(":%d", b.Config.Port), nil)
	for update := range updates {
		log.Printf("update: %+v\n", update)
	}
}
