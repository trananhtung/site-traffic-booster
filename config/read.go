package config

import (
	"log"

	"github.com/spf13/viper"
)

type ProxySource struct {
	Http   []string `mapstructure:"http"`
	Https  []string `mapstructure:"https"`
	Socks4 []string `mapstructure:"socks4"`
	Socks5 []string `mapstructure:"socks5"`
}

const (
	PROXY_SOURCE = "proxy_source"
)

// Read reads the configuration from a JSON file.
func Read() *ProxySource {
	viper.AddConfigPath(".")
	viper.SetConfigType("json")
	viper.SetConfigName(PROXY_SOURCE)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	var t ProxySource
	err = viper.Unmarshal(&t)

	if err != nil {
		log.Fatal(err)
	}
	return &t
}
