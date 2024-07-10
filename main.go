package main

import (
	"encoding/json"
	"log"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	botToken := "7312956632:AAGc_wP8qjEeJvByBV3falkqPdZ_hd9wxFY"
	webhookURL := "https://cafe-ro.com/bot/"

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	webhook, err := tgbotapi.NewWebhook(webhookURL + bot.Token)
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

	http.HandleFunc("/"+bot.Token, func(w http.ResponseWriter, r *http.Request) {
		var update tgbotapi.Update
		if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
			log.Printf("Failed to decode update: %v", err)
			http.Error(w, "Failed to decode update", http.StatusBadRequest)
			return
		}

		// Process the update based on its type (e.g., message, callback query)
		if update.Message != nil {
			log.Printf("Received message: %s", update.Message.Text)

			// Example: Echo the received message back to the sender
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			if _, err := bot.Send(msg); err != nil {
				log.Printf("Error sending message: %v", err)
				http.Error(w, "Failed to send message", http.StatusInternalServerError)
				return
			}
		}

		// Respond to Telegram with a 200 OK status
		w.WriteHeader(http.StatusOK)
	})
	http.ListenAndServe(":4000", nil)

}

// func main() {
// 	botToken := "7312956632:AAGc_wP8qjEeJvByBV3falkqPdZ_hd9wxFY"

// 	bot, err := tgbotapi.NewBotAPI(botToken)
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	bot.Debug = true

// 	log.Printf("Authorized on account %s", bot.Self.UserName)

// 	u := tgbotapi.NewUpdate(0)
// 	u.Timeout = 60

// 	updates := bot.GetUpdatesChan(u)

// 	for update := range updates {
// 		if update.Message != nil {
// 			switch update.Message.Command() {
// 			case "start":
// 				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Welcome! Use the custom keyboard to navigate.")

// 				keyboard := tgbotapi.NewReplyKeyboard(
// 					tgbotapi.NewKeyboardButtonRow(
// 						tgbotapi.NewKeyboardButton("Option 1"),
// 						tgbotapi.NewKeyboardButton("Option 2"),
// 					),
// 				)

// 				msg.ReplyMarkup = keyboard

// 				bot.Send(msg)
// 			case "help":
// 				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Available commands:\n/start - start the bot\n/help - see this help message")
// 				bot.Send(msg)
// 			default:
// 				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "I don't know that command.")
// 				bot.Send(msg)
// 			}
// 		}
// 	}
// }

// if update.Message != nil {
// 	switch update.Message.Command() {
// 	case "start":
// 		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Welcome to the bot! Type /help to see available commands.")
// 		bot.Send(msg)
// 	case "help":
// 		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Available commands:\n/start - start the bot\n/help - see this help message")
// 		bot.Send(msg)
// 	default:
// 		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "I don't know that command.")
// 		bot.Send(msg)
// 	}
// }

// user auth

// var authorizedUsers = map[int64]bool{
//     12345678: true, // Replace with actual user IDs
//     87654321: true,
// }
//
// for update := range updates {
// 	if update.Message != nil {
// 		if _, authorized := authorizedUsers[update.Message.From.ID]; !authorized {
// 			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You are not authorized to use this bot.")
// 			bot.Send(msg)
// 			continue
// 		}

// 		// Process authorized user's message
// 		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Welcome, authorized user!")
// 		bot.Send(msg)
// 	}
// }
