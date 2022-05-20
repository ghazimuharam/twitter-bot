package app

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/ghazimuharam/twitter-bot/cmd/config/entity"
)

func Init(configs *entity.Config) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	c := InitCronApp(configs)

	go c.StartCron()

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}
