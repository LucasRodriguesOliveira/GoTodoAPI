package config

import "github.com/spf13/viper"

func serverEnv() {
	APP["host"] = viper.GetString("HOST")
	APP["port"] = viper.GetString("PORT")
}

func init() {
	loadList = append(loadList, serverEnv)
}
