package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	botToken := "7312956632:AAGc_wP8qjEeJvByBV3falkqPdZ_hd9wxFY"

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			switch update.Message.Command() {
			case "start":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Welcome! Use the custom keyboard to navigate.")

				keyboard := tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Option 1"),
						tgbotapi.NewKeyboardButton("Option 2"),
					),
				)

				msg.ReplyMarkup = keyboard

				bot.Send(msg)
			case "help":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Available commands:\n/start - start the bot\n/help - see this help message")
				bot.Send(msg)
			default:
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "I don't know that command.")
				bot.Send(msg)
			}
		}
	}
}

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
