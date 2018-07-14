// Copyright 2013 @atotto. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package clipboard read/write on clipboard
package clipboard

import (
	"time"
)

// ReadAll read string from clipboard
func ReadAll() (string, error) {
	return readAll()
}

// WriteAll write string to clipboard
func WriteAll(text string) error {
	return writeAll(text)
}

// Monitor starts monitoring the clipboard for changes. When
// a change is detected, it is sent over the channel.
func Monitor(interval time.Duration, stopCh <-chan struct{}, changes chan<- string) error {
	defer close(changes)

	currentValue, err := ReadAll()
	if err != nil {
		return err
	}

	for {
		select {
		case <-stopCh:
			return nil
		default:
			newValue, _ := ReadAll()
			if newValue != currentValue {
				currentValue = newValue
				changes <- currentValue
			}
		}
		time.Sleep(interval)
	}
}

// Unsupported might be set true during clipboard init, to help callers decide
// whether or not to offer clipboard options.
var Unsupported bool
