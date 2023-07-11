package config

import "github.com/spf13/viper"

func InitConfig() {
	viper.AutomaticEnv()
	apiConfig()
	mongoConfig()
	loggerConfig()
}
