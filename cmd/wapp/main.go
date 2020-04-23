package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/zserge/webview"
)

var web webview.WebView
var appSlice []string

var (
	h bool
	v bool
	l bool
)

func main() {
	initAppSlice()
	initFlag()

	parseArgs()

	if web != nil {
		web.Run()
	}
}

func initAppSlice() {
	appSlice = append(appSlice, "wechat", "youdao", "time")
}

func initFlag() {
	flag.BoolVar(&h, "h", false, "help")
	flag.BoolVar(&v, "v", false, "version")
	flag.BoolVar(&l, "l", false, "support app")
	flag.Parse()
	flag.Usage = usage
}

func parseArgs() {
	if h {
		flag.Usage()
		return
	}

	if v {
		fmt.Println("0.0.1")
		return
	}

	if l {
		fmt.Println("support:")
		for _, value := range appSlice {
			fmt.Printf("\t%s\n", value)
		}
		return
	}

	if flag.NArg() == 0 {
		usage()
		return
	}

	apps := flag.Args()
	if len(apps) > 0 {
		switch apps[0] {
		case "wechat", "wx":
			web = makeWebview("wechat", 1000, 800,
				"https://wx.qq.com/", false)
		case "youdao", "yd":
			web = makeWebview("youdao", 1000, 800,
				"https://note.youdao.com/web/", true)
		case "time":
			path, err := filepath.Abs("web/static/time.html")
			if err != nil {
				fmt.Println("Error: web/static/time.html not found")
				os.Exit(0)
			}
			web = makeWebview("time", 600, 600, "file://"+path, true)
		default:
			fmt.Println("No support:", apps[0])
		}
	}
}

func makeWebview(title string, wide int, high int,
	url string, chenge bool) webview.WebView {
	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle(title)
	if chenge {
		w.SetSize(wide, high, webview.HintNone)
	} else {
		w.SetSize(wide, high, webview.HintFixed)
	}
	w.Navigate(url)
	return w
}

func usage() {
	fmt.Println(`example:
    web-app wechat`)
	flag.PrintDefaults()
}
