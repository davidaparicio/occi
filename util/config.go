package util

import (
	"github.com/spf13/viper"
)

type Config struct {
	Endpoint           string `mapstructure:"ENDPOINT"`
	Application_Key    string `mapstructure:"APPLICATION_KEY"`
	Application_Secret string `mapstructure:"APPLICATION_SECRET"`
	Consumer_Key       string `mapstructure:"CONSUMER_KEY"`
	//Sleep_Time_Millisecond int    `mapstructure:"SLEEP_TIME_MS"`
}

// LoadConfig reads configuration from file or environment variables.
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
