// Copyright 2013 @atotto. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package clipboard_test

import (
	. "."
	"testing"
)

func TestCopyAndPaste(t *testing.T) {
	expected := "日本語"

	err := WriteAll(expected)
	if err != nil {
		t.Fatal(err)
	}

	actual, err := ReadAll()
	if err != nil {
		t.Fatal(err)
	}

	if actual != expected {
		t.Errorf("want %s, got %s", expected, actual)
	}
}

func BenchmarkReadAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadAll()
	}
}

func BenchmarkWriteAll(b *testing.B) {
	text := "いろはにほへと"
	for i := 0; i < b.N; i++ {
		WriteAll(text)
	}
}
