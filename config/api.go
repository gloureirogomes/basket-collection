package config

import "github.com/spf13/viper"

func apiConfig() {
	viper.SetDefault("API_PORT", "8888")
}
