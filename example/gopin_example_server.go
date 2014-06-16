package main

import (
	a "github.com/zerklabs/auburn-http"
	"os"
)

func main() {
	ch := make(chan os.Signal)

	go trustedServer()
	go compromisedServer()

	for {
		select {
		case <-ch:
			os.Exit(0)
		}
	}
}

func trustedServer() {
	server := a.New("127.0.0.1:8081")

	server.Options.EnableTLS("certs/example_server.crt", "certs/example_server.key")
	server.AddRouteForMethod("/get", a.GET, getHandler)

	server.Start()
}

func compromisedServer() {
	server := a.New("127.0.0.1:8082")

	server.Options.EnableTLS("certs/mitm_server.crt", "certs/mitm_server.key")
	server.AddRouteForMethod("/get", a.GET, getHandler)
	server.Start()
}

func getHandler(req a.HttpTransaction) {
	req.RespondWithText("Hello, World!")
}
