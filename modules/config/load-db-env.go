package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func LoadDBEnv() map[string]string {
	var (
		host string = viper.GetString("DB_HOST")
		port string = viper.GetString("DB_PORT")
		name string = viper.GetString("DB_NAME")
		pass string = viper.GetString("DB_PASS")
		user string = viper.GetString("DB_USER")
	)

	var dsn string = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		host, user, pass, name, port,
	)

	return map[string]string{
		"host": host,
		"port": port,
		"name": name,
		"pass": pass,
		"user": user,
		"dsn":  dsn,
	}
}
