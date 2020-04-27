package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

var (
	h bool
)

func main() {
	initFlag()
	parseArgs()

}

func parseArgs() {
	if h {
		flag.Usage()
	}

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("dev...")
		os.Exit(0)
	}

	if len(args) >= 2 {
		switch args[0] {
		case "json":
			var str bytes.Buffer
			fmt.Println(args[1])
			err := json.Indent(&str, []byte(args[1]), "", "    ")
			if err != nil {
				fmt.Printf("json indent error %s \n", err)
				return
			}
			fmt.Printf("formated: \n%s\n", str.String())
		default:
			strType := args[0]
			if len(strType) > 10 {
				strType = strType[:10] + "..."
			}

			fmt.Println("No support:", strType)
		}
	} else {
		fmt.Printf("args error: %d\n", len(args))
	}
}

func initFlag() {
	flag.BoolVar(&h, "h", false, "help")

	flag.Parse()
	flag.Usage = usage
}

func usage() {
	fmt.Println(`help: code format demo`)
	flag.PrintDefaults()
}
