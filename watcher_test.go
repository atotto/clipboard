package clipboard

import (
	"fmt"
	"testing"
	"time"
)

func TestChange(t *testing.T) {
	testStrings := []string{
		"content 1",
		"content 2",
		"content 3",
	}

	WriteAll("")
	ch, cancel := Watch(20 * time.Millisecond)

	go func() {
		for _, c := range testStrings {
			WriteAll(c)
			time.Sleep(100 * time.Millisecond)
		}

		// check if rewriting the same content doesn't affect the test
		WriteAll("content 3")
		time.Sleep(100 * time.Millisecond)

		cancel()
	}()

	var i int
	for c := range ch {
		fmt.Println(i, testStrings[i], c)
		if c != testStrings[i] {
			t.Errorf("want %s, got %s", testStrings[i], c)
		}
		i++
	}

	if i != 3 {
		t.Errorf("want %d, got %d", 3, i)
	}
}
