// This assumes the input is already just the subjectPublicKeyInfo
// from whatever certificate you are wanting to pin against
//
// If you only have a certificate and want to get the subjectPublicKeyInfo
// see: get_subject_pubkeyinfo.go
//
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		panic(err)
	}

	formattedTypeBlock := "var trustedPublicKey = []byte{\n"

	counter := 0
	for _, v := range bytes {
		if counter < 8 {
			if counter == 0 {
				formattedTypeBlock += fmt.Sprintf("\t0x%x,", v)
			} else {
				formattedTypeBlock += fmt.Sprintf("0x%x,", v)
			}

			counter += 1
		} else {
			formattedTypeBlock += "\n"
			counter = 0
		}
	}

	formattedTypeBlock += "\n}"

	fmt.Println(formattedTypeBlock)
}
