[![Build Status](https://travis-ci.org/atotto/clipboard.svg?branch=master)](https://travis-ci.org/atotto/clipboard)

[![GoDoc](https://godoc.org/github.com/atotto/clipboard?status.svg)](http://godoc.org/github.com/atotto/clipboard)

# Clipboard for Go

Provide functionality for copying from and pasting to the clipboard.

### Build

    $ go get -u github.com/atotto/clipboard

### Platforms

* macOS
* Windows 7 (probably works on later editions too)
* Linux, Unix (requires `xclip` and `xsel` to be installed (or `wlpaste` and `wlcopy`, for Wayland)

### Documentation

* http://godoc.org/github.com/atotto/clipboard

### Notes

* For functions that takes or return a string, only UTF-8 is supported

### TODO

- [ ] Clipboard watcher(?)

## Commands

paste shell command:

    $ go get -u github.com/atotto/clipboard/cmd/gopaste
    $ # example:
    $ gopaste > document.txt

copy shell command:

    $ go get -u github.com/atotto/clipboard/cmd/gocopy
    $ # example:
    $ cat document.txt | gocopy
