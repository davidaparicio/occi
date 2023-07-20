package util

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/ovh/go-ovh/ovh"
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

// PrintIfNotEmpty prints if not empty, an interface array
func PrintIfNotEmpty(name string, objects []interface{}) {
	if len(objects) == 0 {
		return
	}
	var b strings.Builder
	b.WriteString(name)
	b.WriteString("\t")
	for _, obj := range objects {
		fmt.Fprintf(&b, "%s\t", obj)
	}
	fmt.Println(b.String())
}

// PrintIfNotEmpty prints if not empty, a string array
func PrintIfNotEmptyS(name string, s string) {
	if len(s) == 0 {
		return
	}
	fmt.Println(s)
}

// PrintAssets prints all available services of a given API URL
func PrintAssets(url, name string, client *ovh.Client) {
	var assets []interface{}
	err := client.Get(url, &assets)
	if err != nil {
		log.Fatal("client.Get("+url+"):", err)
	}
	PrintIfNotEmpty(name+":", assets)
	time.Sleep(100 * time.Millisecond)
}
