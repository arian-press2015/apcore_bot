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
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Choose an option:")
			var keyboard = tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Option 1", "opt1"),
					tgbotapi.NewInlineKeyboardButtonData("Option 2", "opt2"),
				),
			)
			msg.ReplyMarkup = keyboard
			bot.Send(msg)
		} else if update.CallbackQuery != nil {
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
				log.Println(err)
			}

			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You chose: "+update.CallbackQuery.Data)
			if _, err := bot.Send(msg); err != nil {
				log.Println(err)
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
