package config

import "github.com/spf13/viper"

var (
	DSN string
	APP map[string]string

	loadList []func()
)

func Load() {
	viper.AutomaticEnv()
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	for _, load := range loadList {
		load()
	}
}
