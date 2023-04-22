package config

import (
	"strconv"

	"github.com/spf13/viper"
)

func ServerHost() string {
	return viper.GetString("HOST")
}

func ServerPort() (int, error) {
	return strconv.Atoi(viper.GetString("PORT"))
}