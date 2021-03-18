package clipboard

import (
	"context"
	"time"
)

// Watch returns a channel for new clipboard content whenever it changes.
// It will return current clipboard content if not empty when initialized.
// Interval is the time between clipboard checks. Take into account that clickboard
// checks need a few milliseconds as well.
func Watch(interval time.Duration) (<-chan string, context.CancelFunc) {
	var (
		ch          = make(chan string)
		lastContent string
		ctx, cancel = context.WithCancel(context.Background())
	)

	go func() {
		defer close(ch)

		for {
			time.Sleep(interval)

			select {
			case <-ctx.Done():
				return
			default:
			}

			content, err := ReadAll()
			if err != nil || content == "" || content == lastContent {
				continue
			}

			lastContent = content
			ch <- content
		}
	}()

	return ch, cancel
}
