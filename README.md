gopin - Certificate Pinning in Go
=================================

## Introduction
This is a proof of concept for enabling certificate pinning in Go. It currently makes an out-of-band call to retrieve the remote host certificate and compares it against a trusted DER encoded public key that you provide. 

See the examples directory for a server and client setup.

## Usage

```
package main

import (
  "fmt"
  "github.com/zerklabs/gopin"
)

func main() {
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
```


## More information:
* [Certificate Pinning][owasp]


[owasp]: https://www.owasp.org/index.php/Certificate_and_Public_Key_Pinning
