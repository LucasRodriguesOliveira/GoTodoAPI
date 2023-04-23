package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func dbEnv() {
	var (
		host string = viper.GetString("DB_HOST")
		port string = viper.GetString("DB_PORT")
		name string = viper.GetString("DB_NAME")
		pass string = viper.GetString("DB_PASS")
		user string = viper.GetString("DB_USER")
	)

	DSN = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		host, user, pass, name, port,
	)
}

func init() {
	loadList = append(loadList, dbEnv)
}
