package config

import "github.com/spf13/viper"

func mongoConfig() {
	viper.SetDefault("MONGO_USER", "basket")
	viper.SetDefault("MONGO_PASSWORD", "password-here")
	viper.SetDefault("MONGO_DATABASE_NAME", "basket-collection")
	viper.SetDefault("MONGO_TEAM_COLLECTION", "team")
}
