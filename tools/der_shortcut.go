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
