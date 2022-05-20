package main

import (
	"github.com/ghazimuharam/twitter-bot/cmd/config"
	"github.com/ghazimuharam/twitter-bot/internal/app"
)

func main() {
	appName := "twitter-bot"

	configs := config.InitConfig(appName)

	app.Init(configs)
}
