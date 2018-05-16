package main

import (
	"flag"
	"fmt"
	"github.com/kiennq/clipboard"
	"os"
	"encoding/json"
	"io"
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
	htmlVal := flag.String("html", "", "html format text")
	textVal := flag.String("text", "", "normal text")
	useStdin := flag.Bool("in", false, "using stdin as input, the input format is json {text:, htmlText:}")
	doClear := flag.Bool("clear", true, "clear previous clipboard")

	flag.Parse()

	htmlText := *htmlVal
	text := *textVal
	if *useStdin {
		type Message struct {
			Text, HtmlText string
		}

		dec := json.NewDecoder(os.Stdin)
		var m Message
		if err := dec.Decode(&m); err != nil && err != io.EOF {
			panic(err)
		}

		htmlText = m.HtmlText
		text = m.Text
	}

	htmlBlock := getHtmlBlock(htmlText)

	if *doClear {
		if err := clipboard.ClearClipboard(); err != nil {
			panic(err)
		}
	}

	if htmlText != "" {
		if err := clipboard.WriteAllWithFormat(htmlBlock, cfHtml); err != nil {
			panic(err)
		}
	}

	if text != "" {
		if err := clipboard.WriteAllWithFormat(text, cfText); err != nil {
			panic(err)
		}
	}
}
