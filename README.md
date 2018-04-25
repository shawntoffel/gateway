# gateway
Simple API gateway using a reverse proxy

    g := gateway.NewGateway()
    
    err := g.Handle("http://localhost:3000/some/endpoint")

    if err != nil {
        return err
    }

    err = g.Start(8080)

    if err != nil {
        return err
    }
