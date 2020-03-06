// Copyright 2013 @atotto. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package clipboard read/write on clipboard
package clipboard

// ReadAll will read a string from the clipboard
func ReadAll() (string, error) {
	return readAll()
}

// WriteAll will write a string to the clipboard
func WriteAll(text string) error {
	return writeAll(text)
}

// ReadAllBytes will read bytes from the clipboard
func ReadAllBytes() ([]byte, error) {
	return readAllBytes()
}

// WriteAllBytes will write bytes to the clipboard
func WriteAllBytes(b []byte) error {
	return writeAllBytes(b)
}

// Unsupported might be set true during clipboard init, to help callers decide
// whether or not to offer clipboard options.
var Unsupported bool
