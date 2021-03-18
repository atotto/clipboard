package main

import (
	"log"
	"time"

	"github.com/spy16/clipboard"
)

func main() {
	changes := make(chan string, 10)
	stopCh := make(chan struct{})

	go clipboard.Monitor(time.Second, stopCh, changes)

	// Watch for changes
	for {
		select {
		case <-stopCh:
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
