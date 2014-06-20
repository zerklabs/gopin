gopin - Certificate Pinning in Go
=================================

## Introduction
This is a proof of concept for enabling certificate pinning in Go. It does this by providing an http.Transport object
with a custom Dial method. If the pin check fails, the Dial method will never work and will only return an error.


## Usage
If the pinned certificate matches the host certificate, it will not return an error and will
return a valid http.Transport. However, if the pinning does fail, any subsequent connection attempt using
the http client with the failed transport will fail. See examples/gopin_example.go


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

  transport, err := gopin.New(ourTrustedPublicKey)

  if err != nil {
    panic(err)
  }

  client := &http.Client{Transport: transport}
  resp, err := client.Get("https://127.0.0.1:8081")

  // do something with your results!
}
```

## Tools

### der_shortcut
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

### get_pkp
Reads a file via STDIN (or via --flag) and outputs a helpful public key fingerprint which can be used in Chrome under chrome://net-internals/#hsts



## More information:
* [Certificate Pinning][owasp]


[owasp]: https://www.owasp.org/index.php/Certificate_and_Public_Key_Pinning
