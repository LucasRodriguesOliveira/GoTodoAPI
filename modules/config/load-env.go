package config

import "github.com/spf13/viper"

func LoadEnv() {
	viper.AutomaticEnv()
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
}
