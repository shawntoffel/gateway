package main

import (
	"github.com/BurntSushi/toml"
	"github.com/shawntoffel/gateway"
	"github.com/shawntoffel/services-core/command"
)

type config struct {
	Port         string
	Destinations []string
}

var args command.CommandArgs

func init() {
	args = command.ParseArgs()
}

func main() {
	c := config{}

	_, err := toml.DecodeFile(args.ConfigFile, &c)

	if err != nil {
		panic(err)
	}

	g := gateway.NewGateway()

	for _, destination := range c.Destinations {
		g.Handle(destination)
	}

	g.Start(c.Port)
}
