// Copyright 2013 @atotto. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package clipboard

import (
	"reflect"
	"testing"
)

func TestPrimaryCommandDoesNotChangeDefaultSelection(t *testing.T) {
	args := []string{"synthetic-clipboard", "--input", "--clipboard"}
	primary := commandForSelection(args, true)
	if want := []string{"synthetic-clipboard"}; !reflect.DeepEqual(primary.Args, want) {
		t.Fatalf("primary command args = %v, want %v", primary.Args, want)
	}

	defaultSelection := commandForSelection(args, false)
	if want := []string{"synthetic-clipboard", "--input", "--clipboard"}; !reflect.DeepEqual(defaultSelection.Args, want) {
		t.Fatalf("default command args after primary selection = %v, want %v", defaultSelection.Args, want)
	}
	if want := []string{"synthetic-clipboard", "--input", "--clipboard"}; !reflect.DeepEqual(args, want) {
		t.Fatalf("original command args = %v, want unchanged %v", args, want)
	}
}
