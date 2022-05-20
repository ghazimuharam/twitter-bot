//go:build wireinject
// +build wireinject

package app

// The build tag makes sure the stub is not built in the final build.
import (
	"github.com/ghazimuharam/twitter-bot/cmd/config/entity"
	"github.com/ghazimuharam/twitter-bot/internal/twitter/delivery/cron"
	"github.com/google/wire"
)

func InitCronApp(configs *entity.Config) *cron.CronApplication {
	wire.Build(allSet)
	return &cron.CronApplication{}
}
