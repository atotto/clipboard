// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin

package clipboard

import (
	"os/exec"
)

type Watcher struct {
	Copied chan string
}

var (
	pbpasteCmd = exec.Command("pbpaste")
	pbcopyCmd  = exec.Command("pbcopy")
)

func ReadAll() string {
	out, err := pbpasteCmd.Output()
	if err != nil {
		panic(err)
	}
	return string(out)
}

func WriteAll(str string) {
	in, err := pbcopyCmd.StdinPipe()
	if err != nil {
		panic(err)
	}

	pbcopyCmd.Start()
	in.Write([]byte(str))
	in.Close()
	pbcopyCmd.Wait()
}
