package gateway

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Gateway interface {
	Handle(destination string) error
	Start(port string) error
}

type gateway struct {
	ErrorLog *log.Logger
	mux      *http.ServeMux
}

func NewGateway() Gateway {
	return NewGatewayWithErrorLog(&log.Logger{})
}

func NewGatewayWithErrorLog(errorLog *log.Logger) Gateway {
	return &gateway{
		mux:      http.NewServeMux(),
		ErrorLog: errorLog,
	}
}

func (g *gateway) Handle(destination string) error {
	destinationUrl, err := url.Parse(destination)

	if err != nil {
		return err
	}

	g.mux.Handle(destinationUrl.Path, g.proxy(destinationUrl))

	return nil
}

func (g *gateway) Start(port string) error {
	server := http.Server{
		Addr:     ":" + port,
		ErrorLog: g.ErrorLog,
	}

	http.Handle("/", g.mux)

	return server.ListenAndServe()
}

func (g *gateway) proxy(destinationUrl *url.URL) *httputil.ReverseProxy {
	destinationHost := &url.URL{
		Scheme: destinationUrl.Scheme,
		Host:   destinationUrl.Host,
	}

	proxy := httputil.NewSingleHostReverseProxy(destinationHost)
	proxy.ErrorLog = g.ErrorLog

	return proxy
}
