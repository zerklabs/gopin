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

## Tools

In the tools/ directory are a few utilities to make your life a little easier:

### der_shortcut.go
Reads a file via STDIN and outputs a go code block for use in your program.

```
$ cat certs/example_server.der | go run der_shortcut.go

var trustedPublicKey = []byte{
  0x30,0x82,0x1,0x22,0x30,0xd,0x6,0x9,
  0x86,0x48,0x86,0xf7,0xd,0x1,0x1,0x1,
  0x0,0x3,0x82,0x1,0xf,0x0,0x30,0x82,
  0xa,0x2,0x82,0x1,0x1,0x0,0xae,0x1b,
  0x50,0xd9,0x85,0xaa,0x48,0xa2,0x8a,0x27,
[...]
```


## More information:
* [Certificate Pinning][owasp]


[owasp]: https://www.owasp.org/index.php/Certificate_and_Public_Key_Pinning
