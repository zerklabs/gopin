package main

import (
	"bytes"
	"crypto/tls"
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

	tlsConfig := &tls.Config{InsecureSkipVerify: true}
	dialer := gopin.New(ourTrustedPublicKey, tlsConfig)
	transport := &http.Transport{
		Dial: dialer.Dial,
	}

	client := &http.Client{Transport: transport}

	fmt.Println("** Showing a successful example (keys match)")
	resp, err := client.Get("https://127.0.0.1:8081/get")

	if err != nil {
		panic(err)
	}

	b := bytes.NewBuffer(nil)
	b.ReadFrom(resp.Body)

	fmt.Println(b.String())

	fmt.Println("** Showing a failed example (keys do not match)")
	resp, err = client.Get("https://127.0.0.1:8082/get")

	if err != nil {
		fmt.Println("Failed to match trusted public key!")
		fmt.Println(err)
	}

}
