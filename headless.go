// Copyright 2015 @atotto. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// You can use the build tag `headless` for unsupported environment.
//
//    $ go build -tags headless
//

// +build headless

package clipboard

func init() {
	Unsupported = true
}
