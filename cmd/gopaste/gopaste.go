package main

import (
	"github.com/atotto/clipboard"

	"fmt"
)

func main() {
	fmt.Print(clipboard.ReadAll())
}
