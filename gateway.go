package gateway

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// Gateway is a simple API gateway
type Gateway interface {
	Handle(destination string) error
	Start(port string) error
}

type gateway struct {
	ErrorLog *log.Logger
	mux      *http.ServeMux
}

// NewGateway creates a new default Gateway
func NewGateway() Gateway {
	return NewGatewayWithErrorLog(&log.Logger{})
}

// NewGatewayWithErrorLog creates a new default Gateway with a custom ErrorLog
func NewGatewayWithErrorLog(errorLog *log.Logger) Gateway {
	return &gateway{
		mux:      http.NewServeMux(),
		ErrorLog: errorLog,
	}
}

//Handle Adds a handler for proxying requests to the provided destination
func (g *gateway) Handle(destination string) error {
	destinationUrl, err := url.Parse(destination)

	if err != nil {
		return err
	}

	g.mux.Handle(destinationUrl.Path, g.proxy(destinationUrl))

	return nil
}

//Start starts the reverse proxy
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
