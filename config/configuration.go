package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Configuration struct {
	ENV []Settings `mapstructure:"env"`
}

type Settings struct {
	NAME  string
	VALUE string
}

func InitConfig(configPath string) {

	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AddConfigPath(configPath)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("error reading config: %s", err.Error())
	}

	var env Configuration
	viper.Unmarshal(&env)
	for _, value := range env.ENV {
		if _, ok := os.LookupEnv(value.NAME); !ok {
			os.Setenv(value.NAME, value.VALUE)
		}

	}
}
