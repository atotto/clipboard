// Copyright 2013 @atotto. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package clipboard

import (
	"syscall"
	"time"
	"unsafe"
)

const (
	cfUnicodetext = 13
	cfText        = 1
	gmemMoveable  = 0x0002
)

var (
	user32                  = syscall.MustLoadDLL("user32")
	openClipboard           = user32.MustFindProc("OpenClipboard")
	closeClipboard          = user32.MustFindProc("CloseClipboard")
	emptyClipboard          = user32.MustFindProc("EmptyClipboard")
	getClipboardData        = user32.MustFindProc("GetClipboardData")
	setClipboardData        = user32.MustFindProc("SetClipboardData")
	registerClipboardFormat = user32.MustFindProc("RegisterClipboardFormatW")

	kernel32     = syscall.NewLazyDLL("kernel32")
	globalAlloc  = kernel32.NewProc("GlobalAlloc")
	globalFree   = kernel32.NewProc("GlobalFree")
	globalSize   = kernel32.NewProc("GlobalSize")
	globalLock   = kernel32.NewProc("GlobalLock")
	globalUnlock = kernel32.NewProc("GlobalUnlock")
	memcpy       = kernel32.NewProc("RtlCopyMemory")
)

// waitOpenClipboard opens the clipboard, waiting for up to a second to do so.
func waitOpenClipboard() error {
	started := time.Now()
	limit := started.Add(time.Second)
	var r uintptr
	var err error
	for time.Now().Before(limit) {
		r, _, err = openClipboard.Call(0)
		if r != 0 {
			return nil
		}
		time.Sleep(time.Millisecond)
	}
	return err
}

func readAll() (string, error) {
	return readAllWithFormat(cfText)
}

func readAllWithFormat(cf uintptr) (string, error) {
	err := waitOpenClipboard()
	if err != nil {
		return "", err
	}
	defer closeClipboard.Call()

	h, _, err := getClipboardData.Call(cf)
	if h == 0 {
		if err != syscall.Errno(0) {
			return "", err
		}

		return "", nil
	}

	size, _, err := globalSize.Call(h)
	if size == 0 {
		return "", err
	}

	l, _, err := globalLock.Call(h)
	if l == 0 {
		return "", err
	}

	text := string((*[1 << 20]byte)(unsafe.Pointer(l))[:(size - 1)])

	r, _, err := globalUnlock.Call(h)
	if r == 0 {
		return "", err
	}

	return text, nil
}

func writeAll(text string) error {
	return writeAllWithFormat(text, cfText)
}

func writeAllWithFormat(text string, cf uintptr) error {
	err := waitOpenClipboard()
	if err != nil {
		return err
	}
	defer closeClipboard.Call()

	data := syscall.StringByteSlice(text)

	// "If the hMem parameter identifies a memory object, the object must have
	// been allocated using the function with the GMEM_MOVEABLE flag."
	h, _, err := globalAlloc.Call(gmemMoveable, uintptr(len(data)))
	if h == 0 {
		return err
	}

	defer func() {
		if h != 0 {
			globalFree.Call(h)
		}
	}()

	l, _, err := globalLock.Call(h)
	if l == 0 {
		return err
	}

	r, _, err := memcpy.Call(l, uintptr(unsafe.Pointer(&data[0])), uintptr(len(data)))
	if r == 0 {
		return err
	}

	r, _, err = globalUnlock.Call(h)
	if r == 0 {
		if err.(syscall.Errno) != 0 {
			return err
		}
	}

	r, _, err = setClipboardData.Call(cf, h)
	if r == 0 {
		return err
	}

	h = 0 // suppress deferred cleanup
	return nil
}

func getClipboardFormat(text string) (uintptr, error) {
	ptr, err := syscall.UTF16PtrFromString(text)
	if err != nil {
		return 0, err
	}

	cf, _, err := registerClipboardFormat.Call(uintptr(unsafe.Pointer(ptr)))
	if cf == 0 {
		return 0, err
	}

	return cf, nil
}

func clearClipboard() error {
	err := waitOpenClipboard()
	if err != nil {
		return err
	}
	defer closeClipboard.Call()

	r, _, err := emptyClipboard.Call()
	if r == 0 {
		return err
	}

	return nil
}
