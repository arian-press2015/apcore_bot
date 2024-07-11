package httpserver

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/arian-press2015/apcore_bot/bot"
	"github.com/arian-press2015/apcore_bot/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type HTTPServer struct {
	Bot    *bot.TelegramBot
	Config *config.Config
}

func NewHTTPServer(bot *bot.TelegramBot, config *config.Config) *HTTPServer {
	return &HTTPServer{Bot: bot, Config: config}
}

func (s *HTTPServer) HandleWebhook(w http.ResponseWriter, r *http.Request) {
	var update tgbotapi.Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		log.Printf("Failed to decode update: %v", err)
		http.Error(w, "Failed to decode update", http.StatusBadRequest)
		return
	}

	if update.Message != nil {
		log.Printf("Received message: %s", update.Message.Text)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		if _, err := s.Bot.Bot.Send(msg); err != nil {
			log.Printf("Error sending message: %v", err)
			http.Error(w, "Failed to send message", http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}
