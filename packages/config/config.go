package config

import (
	"github.com/spf13/viper"
)

var AocSession string

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.ReadInConfig()

	AocSession = viper.GetString("aoc.session")
}
