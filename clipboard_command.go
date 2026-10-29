// Copyright 2013 @atotto. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package clipboard

import "os/exec"

func commandForSelection(args []string, primary bool) *exec.Cmd {
	if primary {
		args = args[:1]
	}
	return exec.Command(args[0], args[1:]...)
}
