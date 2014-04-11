// Copyright 2013 @atotto. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build linux

package clipboard

import (
	"os/exec"
)

var (
	pasteCmdArgs []string
	copyCmdArgs  []string

	xselPaste = []string{"xsel", "--output", "--clipboard"}
	xselCopy  = []string{"xsel", "--input", "--clipboard"}

	xclipPaste = []string{"xclip", "-out", "-secondary", "clipboard"}
	xclipCopy  = []string{"xclip", "-in", "-secondary", "clipboard"}
)

func init() {
	_, err := exec.LookPath("xsel")
	if err == nil {
		pasteCmdArgs = xselPaste
		copyCmdArgs = xselCopy
		return
	}

	_, err = exec.LookPath("xclip")
	if err == nil {
		pasteCmdArgs = xclipPaste
		copyCmdArgs = xclipCopy
		return
	}

	println("No clipboard utilities available. Please install xset or xclip.")
}

func getPasteCommand() *exec.Cmd {
	return exec.Command(pasteCmdArgs[0], pasteCmdArgs[1:]...)
}

func getCopyCommand() *exec.Cmd {
	return exec.Command(copyCmdArgs[0], copyCmdArgs[1:]...)
}

func readAll() (string, error) {
	pasteCmd := getPasteCommand()
	out, err := pasteCmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func writeAll(text string) error {
	copyCmd := getCopyCommand()
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
