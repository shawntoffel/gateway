package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/go-kit/kit/log"
	configreader "github.com/shawntoffel/services-core/config"
	"github.com/shawntoffel/services-core/runner"
)

type Config struct {
	Port         int
	Destinations []Service
}

type Service struct {
	Host     string
	Endpoint string
}

type Gateway interface {
	Start() error
}

type gateway struct {
	logger log.Logger
	config Config
}

func NewGateway(logger log.Logger, config Config) Gateway {
	return &gateway{logger, config}
}

func (g *gateway) Start() error {
	mux := http.NewServeMux()

	for _, service := range g.config.Destinations {
		url, err := url.Parse(service.Host)

		if err != nil {
			return err
		}

		g.logger.Log("destination", url, "endpoint", service.Endpoint)

		mux.Handle(service.Endpoint, httputil.NewSingleHostReverseProxy(url))
	}

	runner.StartService(mux, g.logger, configreader.ServiceConfig{Port: g.config.Port})

	return nil
}
