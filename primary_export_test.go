package clipboard_test

import (
	"testing"

	"github.com/atotto/clipboard"
)

func TestPrimaryIsDefined(t *testing.T) {
	// Must compile on every GOOS; Primary is a no-op off Unix.
	clipboard.Primary = false
	_ = clipboard.Primary
}
