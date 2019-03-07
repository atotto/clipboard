// Copyright 2013 @atotto. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package clipboard read/write on clipboard
package clipboard

// ReadAll read string from clipboard
func ReadAll() (string, error) {
	return readAll()
}

func ReadAllWithFormat(cf uintptr) (string, error) {
	return readAllWithFormat(cf)
}

// WriteAll write string to clipboard
func WriteAll(text string) error {
	return writeAll(text)
}

// WriteAllWithFormat write string to clipboard with format
func WriteAllWithFormat(text string, cf uintptr) error {
	return writeAllWithFormat(text, cf)
}

func ClearClipboard() error {
	return clearClipboard()
}

func GetClipboardFormat(format string) (uintptr, error) {
	return getClipboardFormat(format)
}

// Unsupported might be set true during clipboard init, to help callers decide
// whether or not to offer clipboard options.
var Unsupported bool
