package util

import (
	"github.com/spf13/viper"	
)

type Config struct {
	ELASTICSEARCH_URL string `mapstructure:"ELASTICSEARCH_URL"`
	USERNAME string `mapstructure:"USERNAME"`
	PASSWORD string `mapstructure:"PASSWORD"`
	INDEX_NAME string `mapstructure:"INDEX_NAME"`
}

// declaring config funct so as to access the variables anywhere
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)	
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
	


