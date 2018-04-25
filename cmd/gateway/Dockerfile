FROM golang:latest as build
ADD . /src
WORKDIR /src
RUN go get -d ./... && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/gateway

FROM alpine:latest
WORKDIR /app
COPY --from=build /src/bin/gateway .
ENTRYPOINT ["./gateway"]
