// Copyright 2013 @atotto. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin

package clipboard

import (
	"os/exec"
)

var (
	pbpasteCmd = exec.Command("pbpaste")
	pbcopyCmd  = exec.Command("pbcopy")
)

func readAll() (string, error) {
	out, err := pbpasteCmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func writeAll(text string) error {
	in, err := pbcopyCmd.StdinPipe()
	if err != nil {
		return err
	}

	if err := pbcopyCmd.Start(); err != nil {
		return err
	}
	if _, err := in.Write([]byte(text)); err != nil {
		return err
	}
	if err := in.Close(); err != nil {
		return err
	}
	return pbcopyCmd.Wait()
}
