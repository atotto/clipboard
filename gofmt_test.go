// Copyright 2015 @atotto. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package clipboard_test

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"

	"testing"
)

type FmtError struct {
	Diff   []byte
	Header string
}

func NewFmtError(diff []byte) *FmtError {
	n := bytes.IndexRune(diff, '\n')
	return &FmtError{diff, string(diff[0:n])}
}

func (e *FmtError) Error() string {
	return e.Header
}

func (e *FmtError) Detail() string {
	return string(e.Diff)
}

func checkFmt(dir string) error {
	cmd := exec.Command("gofmt", "-d", dir)
	diff, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	if len(diff) != 0 {
		return NewFmtError(diff)
	}
	return nil
}

func visitDir(path string, f os.FileInfo, err error) error {
	if err == nil && f.IsDir() {
		err = checkFmt(path)
	}
	if err != nil {
		return err
	}
	return nil
}

func walkDir(path string) error {
	return filepath.Walk(path, visitDir)
}

func TestCheckGofmt(t *testing.T) {
	err := walkDir("./")
	if err != nil {
		if err, ok := err.(*FmtError); ok {
			t.Log(err.Detail())
			t.FailNow()
		}
		t.Fatal(err)
	}
}
