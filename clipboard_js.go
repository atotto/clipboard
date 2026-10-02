// Copyright 2013 @atotto. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build js
// +build js

package clipboard

import "errors"

var errJSUnsupported = errors.New("The system clipboard is not reachable from js/wasm through this synchronous interface (the browser exposes only the asynchronous navigator.clipboard API).")

func init() {
	Unsupported = true
}

func readAll() (string, error) {
	return "", errJSUnsupported
}

func writeAll(text string) error {
	return errJSUnsupported
}
