// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package clipboard_test

import (
	. "."
	"testing"
)

func TestCopyAndPaste(t *testing.T) {
	expected := "testtest"
	Set(expected)
	actual := Get()

	if actual != expected {
		t.Errorf("want %s, got %s", expected, actual)
	}
}
