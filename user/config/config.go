package config

import (
	"github.com/spf13/viper"
	"os"
)

func InitConfig() {
	wordDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(wordDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		return
	}
}
