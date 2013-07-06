// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin

package clipboard

import (
	"log"
	"os/exec"
)

func Get() string {
	cmd := exec.Command("pbpaste")
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	return string(out)
}

func Set(str string) {
	cmd := exec.Command("pbcopy")
	in, err := cmd.StdinPipe()
	if err != nil {
		log.Println(err)
		return
	}

	cmd.Start()
	in.Write([]byte(str))
	in.Close()
	cmd.Wait()
}
