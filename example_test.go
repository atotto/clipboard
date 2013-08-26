package clipboard_test

import (
	"github.com/atotto/clipboard"

	"fmt"
)

func Example() {

	clipboard.WriteAll("日本語")
	text, _ := clipboard.ReadAll()
	fmt.Println(text)

	// Output:
	// 日本語
}
