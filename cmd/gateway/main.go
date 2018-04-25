package main

import (
	"flag"
	"strings"

	"github.com/shawntoffel/gateway"
)

type config struct {
	Port         int
	Destinations *destinations
}

type destinations struct {
	Urls []string
}

func (d *destinations) Set(value string) error {
	split := strings.Split(value, ",")
	for _, url := range split {
		d.Urls = append(d.Urls, url)
	}

	return nil
}

func (d *destinations) String() string {
	return strings.Join(d.Urls, ",")
}

var args config

func init() {
	flag.IntVar(&args.Port, "p", 8080, "Port that the gateway will listen on.")

	d := destinations{}

	flag.Var(&d, "d", "Comma delimited list of destination urls.")
	args.Destinations = &d

	flag.Parse()
}

func main() {
	g := gateway.NewGateway()

	for _, destination := range args.Destinations.Urls {
		err := g.Handle(destination)

		if err != nil {
			panic(err)
		}
	}

	panic(g.Start(args.Port))
}
