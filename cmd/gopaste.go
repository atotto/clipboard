package main

import (
	"fmt"
	"github.com/atotto/clipboard"
)

func main() {
	fmt.Println(clipboard.ReadAll())
}
