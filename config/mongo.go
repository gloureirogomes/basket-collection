package config

import "github.com/spf13/viper"

func mongoConfig() {
	viper.SetDefault("MONGO_HOST", "localhost")
	viper.SetDefault("MONGO_PORT", "27017")
	viper.SetDefault("MONGO_USER", "basket")
	viper.SetDefault("MONGO_PASSWORD", "password-here")
	viper.SetDefault("MONGO_DATABASE_NAME", "basket-collection")
	viper.SetDefault("MONGO_PLAYER_COLLECTION", "player")
}
