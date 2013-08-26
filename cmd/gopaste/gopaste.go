package main

import (
	"github.com/atotto/clipboard"

	"fmt"
)

func main() {
	text, err := clipboard.ReadAll()
	if err != nil {
		panic(err)
	}

	fmt.Print(text)
}
