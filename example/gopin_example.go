package main

import (
	"fmt"
	"github.com/zerklabs/gopin"
)

func main() {
	pinningFailure()
	pinningSuccess()
}

func pinningFailure() {
	// This would be the known and trusted public key that you would manually
	// setup. This is only for example purposes
	ourTrustedPublicKey, err := gopin.ReadInTrustedPublicKey("certs/example_server.der")

	if err != nil {
		panic(err)
	}

	transport, err := gopin.New("127.0.0.1:8082", ourTrustedPublicKey.PublicKeyInfo)

	if err != nil {
		panic(err)
	}

	if transport.State == true {
		panic("This wasn't supposed to happen..")
	} else {
		fmt.Println("Failed to match trusted public key!")
	}
}

func pinningSuccess() {
	ourTrustedPublicKey, err := gopin.ReadInTrustedPublicKey("certs/example_server.der")

	if err != nil {
		panic(err)
	}

	transport, err := gopin.New("127.0.0.1:8081", ourTrustedPublicKey.PublicKeyInfo)

	if err != nil {
		panic(err)
	}

	if transport.State == false {
		panic("This wasn't supposed to happen..")
	} else {
		fmt.Println("Remote host certificate matched our trusted public key!")
	}
}
