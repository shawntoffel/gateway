package main

import (
	"github.com/BurntSushi/toml"
	"github.com/shawntoffel/gateway"
	"github.com/shawntoffel/services-core/command"
)

type Config struct {
	Port         string
	Destinations []string
}

var args command.CommandArgs

func init() {
	args = command.ParseArgs()
}

func main() {
	config := Config{}

	_, err := toml.DecodeFile(args.ConfigFile, &config)

	if err != nil {
		panic(err)
	}

	g := gateway.NewGateway()

	for _, destination := range config.Destinations {
		g.Handle(destination)
	}

	g.Start(config.Port)
}
