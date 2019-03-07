package clipboard

import (
	"fmt"
)

func Example() {
	WriteAll("日本語")
	text, _ := ReadAll()
	fmt.Println(text)

	// Output:
	// 日本語
}
