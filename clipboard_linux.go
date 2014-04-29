// Copyright 2013 @atotto. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build linux

package clipboard

import (
	"os/exec"
)

const (
	xsel  = "xsel"
	xclip = "xclip"
)

var (
	pasteCmd, copyCmd *exec.Cmd

	xselPaste = exec.Command(xsel, "--output", "--clipboard")
	xselCopy  = exec.Command(xsel, "--input", "--clipboard")

	xclipPaste = exec.Command(xclip, "-out", "-sel", "clipboard")
	xclipCopy  = exec.Command(xclip, "-in", "-sel", "clipboard")
)

func init() {
	pasteCmd = xselPaste
	copyCmd = xselCopy

	_, err := exec.LookPath(xsel)
	if err == nil {
		return
	}

	pasteCmd = xclipPaste
	copyCmd = xclipCopy

	_, err = exec.LookPath(xclip)
	if err == nil {
		return
	}

	println("No clipboard utilities available. Please install xset or xclip.")
}

func readAll() (string, error) {
	out, err := pasteCmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func writeAll(text string) error {
	in, err := copyCmd.StdinPipe()
	if err != nil {
		return err
	}

	if err := copyCmd.Start(); err != nil {
		return err
	}
	if _, err := in.Write([]byte(text)); err != nil {
		return err
	}
	if err := in.Close(); err != nil {
		return err
	}
	return copyCmd.Wait()
}
