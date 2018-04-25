# gateway
Simple API gateway using a reverse proxy

    g := gateway.NewGateway()
    
    err := g.Handle("http://localhost:3000/some/endpoint")

    if err != nil {
        return err
    }

    // requests to http://localhost:8080/some/endpoint will be proxied to http://localhost:3000/some/endpoint
    err = g.Start(8080)

    if err != nil {
        return err
    }
