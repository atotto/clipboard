package main

import (
	"github.com/atotto/clipboard"

	"io/ioutil"
	"os"
)

func main() {

	out, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	clipboard.WriteAll(string(out))
}
