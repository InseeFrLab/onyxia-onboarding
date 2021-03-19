package main

import (
	_ "embed"
	"log"
	"strings"

	"github.com/spf13/viper"
)

//go:embed config.yaml
var s string
var config configuration

func loadConfiguration() {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigType("yaml")
	viper.ReadConfig(strings.NewReader(s)) // Reading defaults
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
	viper.SetConfigName("config.local")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.MergeInConfig()

	err := viper.Unmarshal(&config)
	log.Println(&config)
	if err != nil {
		panic(err)
	}
}