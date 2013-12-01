[![Build Status](https://drone.io/github.com/atotto/clipboard/status.png)](https://drone.io/github.com/atotto/clipboard/latest)

# Clipboard for Go

Provide copying and pasting to the Clipboard for Go.

Download shell commands at https://drone.io/github.com/atotto/clipboard/files

Build:

    $ go get github.com/atotto/clipboard

platform, works on:

* OSX
* Windows 7 (probably work on other Windows)


Document: 

* http://godoc.org/github.com/atotto/clipboard

Notes:

* Text string only
* UTF-8 text encoding only (no conversion)

TODO:

* Linux support
* BSD support
* Clipboard watcher(?)

## Commands:

paste shell command:

    $ go get github.com/atotto/clipboard/cmd/gopaste
    $ # example:
    $ gopaste > document.txt

copy shell command:

    $ go get github.com/atotto/clipboard/cmd/gocopy
    $ # example:
    $ cat document.txt | gocopy



