package config

import (
	"github.com/ak-karimzai/cp-db/internal/logger"
	"github.com/spf13/viper"
)

func InitConfig(fileName string) *viper.Viper {
	config := viper.New()

	config.SetConfigFile(fileName)

	config.AddConfigPath(".")
	config.AddConfigPath("$HOME")

	err := config.ReadInConfig()
	if err != nil {
		logger.GetLogger().Fatalf("Error while parsing configuration file ", err)
	}
	return config
}
