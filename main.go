package main

import (
	"os"

	"github.com/shawntoffel/services-core/command"
	"github.com/shawntoffel/services-core/logging"
	"github.com/spf13/viper"
)

var args command.CommandArgs

func init() {
	args = command.ParseArgs()
}

func main() {
	logger := logging.GetLogger("api-gateway", os.Stdout)

	config := Config{}

	viper.SetConfigFile(args.ConfigFile)
	viper.ReadInConfig()
	viper.Unmarshal(&config)

	gateway := NewGateway(logger, config)

	gateway.Start()
}
