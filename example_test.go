package clipboard_test

import (
	"github.com/atotto/clipboard"

	"fmt"
)

func Example() {

	clipboard.WriteAll("日本語")
	fmt.Println(clipboard.ReadAll())

	// Output:
	// 日本語
}
