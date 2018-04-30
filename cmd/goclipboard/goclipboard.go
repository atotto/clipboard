package main

import (
	"flag"
	"fmt"
	"github.com/kiennq/clipboard"
)

const (
	HtmlMarkerBlock = "Version:0.9\n" +
		"StartHTML:%010d\n" +
		"EndHTML:%010d\n" +
		"StartFragment:%010d\n" +
		"EndFragment:%010d\n"

	HtmlHeader = "<html>\n<body>\n"
	HtmlFooter = "\n</body>\n</html>\n"
)

var (
	cfText uintptr = 1
	cfHtml         = getHtmlClipboard()
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

func main() {
	htmlText := flag.String("html", "", "html format text")
	text := flag.String("text", "", "text")

	flag.Parse()

	htmlBlock := getHtmlBlock(*htmlText)

	if err := clipboard.ClearClipboard(); err != nil {
		panic(err)
	}

	if err := clipboard.WriteAllWithFormat(htmlBlock, cfHtml); err != nil {
		panic(err)
	}

	if err := clipboard.WriteAllWithFormat(*text, cfText); err != nil {
		panic(err)
	}
}
