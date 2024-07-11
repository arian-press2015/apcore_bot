package main

import (
	"context"
	"log"

	"github.com/arian-press2015/apcore_bot/bot"
	"github.com/arian-press2015/apcore_bot/config"
	httpserver "github.com/arian-press2015/apcore_bot/http_server"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			config.NewConfig,
			bot.NewTelegramBot,
			httpserver.NewHTTPServer,
		),
		fx.Invoke(RunBot),
	)

	if err := app.Start(context.Background()); err != nil {
		log.Fatalf("Failed to start application: %v", err)
	}

	<-app.Done()
}

func RunBot(lifecycle fx.Lifecycle, bot *bot.TelegramBot, server *httpserver.HTTPServer) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			bot.Run(server.HandleWebhook)
			return nil
		},
		OnStop: func(context.Context) error {
			// Implement cleanup logic if needed
			return nil
		},
	})
}
