package main

import (
	"bytes"
	"fmt"
	"github.com/zerklabs/gopin"
	"net/http"
)

func main() {
	// This would be the known and trusted public key that you would manually
	// setup. This is only for example purposes
	ourTrustedPublicKey, err := gopin.ReadInTrustedPublicKey("certs/example_server.der")

	if err != nil {
		panic(err)
	}

	transport, err := gopin.New(ourTrustedPublicKey, nil)

	if err != nil {
		panic(err)
	}

	client := &http.Client{
		Transport: transport,
	}

	// this should fail since the public key returned from
	// the server and the pinned key do not match
	_, err = client.Get("https://127.0.0.1:8082/get")

	if err != nil {
		fmt.Println("Failed to match trusted public key!")
		fmt.Println(err)
		fmt.Println("")
	}

	// Now we pass in the this should succeed since the public key returned from
	// the server and the pinned key match
	resp, err := client.Get("https://127.0.0.1:8081/get")

	if err != nil {
		panic(err)
	}

	b := bytes.NewBuffer(nil)

	fmt.Println("Successful pin comparison, should see Hello, World! below")

	// Hello, World!
	b.ReadFrom(resp.Body)
	fmt.Println(b.String())
}
