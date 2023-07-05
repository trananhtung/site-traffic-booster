package main

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/trananhtung/site-traffic-booster/config"
)

var proxySource *config.ProxySource

func init() {
	pflag.StringP("websites", "w", "", "websites to check")
	pflag.Uint16P("user", "u", 5, "number user for each proxy")
	pflag.Uint16P("max", "p", 20, "max threads")
	pflag.Uint16P("timeout", "t", 10, "timeout for each request in seconds")
	pflag.Uint16P("visitTime", "v", 60, "time to visit each website in seconds")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	proxySource = config.Read()
}

func main() {
	fmt.Println(proxySource.Http)
}
