package config

import (
	"log"
	"os"

	"github.com/ghazimuharam/twitter-bot/cmd/config/entity"
	"github.com/spf13/viper"
)

func InitConfig(appName string) *entity.Config {
	env := getEnvironment()

	return readConfigs(appName, env)
}

func readConfigs(appName, env string) *entity.Config {
	configs := &entity.Config{}
	configTypes := []string{"main"}
	for _, configType := range configTypes {
		viper.SetConfigName(appName + "." + configType + "." + env) // name of config file (without extension)
		viper.SetConfigType("json")

		if env == "production" {
			viper.AddConfigPath("files/" + env)
		} else {
			viper.AddConfigPath(env)
		}

		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				log.Fatal("config file not found")
			} else {
				log.Fatal("cannot read config file")
			}
		}
	}


	err := viper.Unmarshal(configs)
	if err != nil {
		log.Fatal("cannot unmarshal viper config")
	}

	return configs
}

func getEnvironment() string {
	env := os.Getenv("ENV")

	if env != "production" {
		return "development"
	}

	return env
}
