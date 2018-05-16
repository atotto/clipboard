package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/kiennq/clipboard"
	"io/ioutil"
	"os"
	"strconv"
)

const (
	HtmlMarkerBlock = "Version:0.9\n" +
		"StartHTML:%010d\n" +
		"EndHTML:%010d\n" +
		"StartFragment:%010d\n" +
		"EndFragment:%010d\n"

	HtmlHeader = "<html>\n<body>\n"
	HtmlFooter = "\n</body>\n</html>"
)

var (
	cfText uintptr = 1
)

func getHtmlBlock(htmlText string) string {
	testPrefix := fmt.Sprintf(HtmlMarkerBlock, 0, 0, 0, 0)
	markerBlockLen := len(testPrefix) // Should be always 100, can be optimized
	headerLen := len(HtmlHeader)
	footerLen := len(HtmlFooter)
	prefix := fmt.Sprintf(HtmlMarkerBlock,
		markerBlockLen,
		markerBlockLen+headerLen+len(htmlText)+footerLen,
		markerBlockLen+headerLen,
		markerBlockLen+headerLen+len(htmlText))
	return prefix + HtmlHeader + htmlText + HtmlFooter
}

func getHtmlClipboard() uintptr {
	cf, err := clipboard.GetClipboardFormat("HTML Format")
	if err != nil {
		panic(err)
	}

	return cf
}

func getCfFromString(format string) uintptr {
	cf := cfText
	if format != "" {
		atoiCf, err := strconv.Atoi(format)
		cf = uintptr(atoiCf)
		if err != nil {
			// not integer, register new clipboard format
			if cf, err = clipboard.GetClipboardFormat(format); err != nil {
				panic(err)
			}
		}
	}

	return cf
}

func doCopy(dontClear bool) {
	if !dontClear {
		if err := clipboard.ClearClipboard(); err != nil {
			panic(err)
		}
	}

	type Message struct {
		Cf, Data string
	}

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	var m []Message
	if err := json.Unmarshal(data, &m); err != nil {
		m = []Message{{"", string(data)}}
	}

	for _, d := range m {
		cf := getCfFromString(d.Cf)
		if d.Data != "" {
			switch d.Cf {
			case "HTML Format":
				d.Data = getHtmlBlock(d.Data)
			}

			if err := clipboard.WriteAllWithFormat(d.Data, cf); err != nil {
				panic(err)
			}
		}
	}
}

func doPaste(format string) {
	cf := getCfFromString(format)
	out, err := clipboard.ReadAllWithFormat(cf)
	if err != nil {
		panic(err)
	}

	fmt.Println(out)
}

func usage() {
	fmt.Println("usage: goclip <command> [<opts>]")
	fmt.Println("Available commands are:")
	fmt.Println("	copy		Copy to clipboard through stdin. The json format is [{cf:, data:},]")
	fmt.Println("	paste		Read from clipboard")
}

func main() {
	copyCmd := flag.NewFlagSet("copy", flag.ExitOnError)
	dontClear := copyCmd.Bool("no-clear", false, "don't clear previous clipboard")

	pasteCmd := flag.NewFlagSet("paste", flag.ExitOnError)
	format := pasteCmd.String("format", "", "paste from clipboard format. Can be string or integer.")

	flag.Usage = usage
	flag.Parse()

	switch flag.Arg(0) {
	case "copy":
		copyCmd.Parse(os.Args[2:])
		doCopy(*dontClear)
	case "paste":
		pasteCmd.Parse(os.Args[2:])
		doPaste(*format)
	default:
		usage()
		os.Exit(2)
	}
}
