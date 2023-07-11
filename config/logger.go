package config

import "github.com/spf13/viper"

func loggerConfig() {
	viper.SetDefault("LOG_LEVEL", "DEBUG")
}
