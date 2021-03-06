# gateway
[![GoDoc](https://godoc.org/github.com/shawntoffel/gateway?status.svg)](https://godoc.org/github.com/shawntoffel/gateway) [![Go Report Card](https://goreportcard.com/badge/github.com/shawntoffel/gateway)](https://goreportcard.com/report/github.com/shawntoffel/gateway) [![Build Status](https://travis-ci.org/shawntoffel/gateway.svg?branch=master)](https://travis-ci.org/shawntoffel/gateway)

Simple API gateway using a reverse proxy

```go
destinations := []string{
    "http://localhost:3000/some/endpoint",
    "http://localhost:3001/some/other/endpoint",
}

g := gateway.NewGateway()
    
for _, destination := range destinations {
    // register an endpoint
    err := g.Handle(destination)

    if err != nil {
        return err
    }
}

// requests to port 8080 will be proxied to the appropriate destination by url path
err = g.Start(8080)

if err != nil {
    return err
}
```
