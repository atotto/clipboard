// Copyright 2013 @atotto. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build linux

package clipboard

import (
	"os/exec"
)

var (
	pasteCmd = exec.Command("xsel --output --clipboard")
	copyCmd  = exec.Command("xsel --input --clipboard")
)

func readAll() string {
	out, err := pasteCmd.Output()
	if err != nil {
		panic(err)
	}
	return string(out)
}

func writeAll(str string) {
	in, err := copyCmd.StdinPipe()
	if err != nil {
		panic(err)
	}

	copyCmd.Start()
	in.Write([]byte(str))
	in.Close()
	copyCmd.Wait()
}
