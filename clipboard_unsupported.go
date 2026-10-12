// Copyright 2013 @atotto. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build js || wasip1 || wasip2

package clipboard

import (
	"errors"
	"runtime"
)

func init() { Unsupported = true }

func readAll() (string, error) {
	return "", errors.New("clipboard: not supported on " + runtime.GOOS)
}

func writeAll(text string) error {
	return errors.New("clipboard: not supported on " + runtime.GOOS)
}
