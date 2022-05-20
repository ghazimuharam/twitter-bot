package cron

import (
	"fmt"
	"log"

	"github.com/ghazimuharam/twitter-bot/cmd/config/entity"
	"github.com/ghazimuharam/twitter-bot/internal/twitter"
	internal_entity "github.com/ghazimuharam/twitter-bot/internal/twitter/entity"
	"github.com/robfig/cron/v3"
)

type CronApplication struct {
	cron     *cron.Cron
	uc       twitter.CronUCitf
	cronList internal_entity.CronList
	configs  *entity.Config
}

func NewCronApplication(uc twitter.CronUCitf, configs *entity.Config) *CronApplication {
	c := cron.New(cron.WithChain(
		cron.Recover(cron.DefaultLogger),
	))
	if configs.App.FeatureFlag.IsCronWithSeconds {
		c = cron.New(cron.WithSeconds(), cron.WithChain(
			cron.Recover(cron.DefaultLogger),
		))
	}

	return &CronApplication{
		cron: c,
		uc:   uc,
		cronList: internal_entity.CronList{
			CronFunction: []internal_entity.CronFunction{
				{
					Function: uc.TweetFromDirectMessage,
					Crontab:  configs.App.CronList.ReadMessage,
				},
			},
		},
	}
}

func (c *CronApplication) StartCron() {
	for _, cron := range c.cronList.CronFunction {
		_, err := c.cron.AddFunc(cron.Crontab, cron.Function)
		if err != nil {
			log.Fatalf("err: %v", err)
		}
	}

	fmt.Println("Cron started")
	c.cron.Start()
}
