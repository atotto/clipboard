package main

import (
	"context"
	"log"
	"time"

	"github.com/shivylp/clipboard"
)

func main() {
	changes := make(chan string, 10)
	ctx := context.Background()

	go clipboard.Monitor(ctx, time.Second, changes)

	// Watch for changes
	for {
		select {
		case <-ctx.Done():
			break
		default:
			change, ok := <-changes
			if ok {
				log.Printf("change received: '%s'", change)
			} else {
				log.Printf("channel has been closed. exiting..")
			}
		}
	}
}
