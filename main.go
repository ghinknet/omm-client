
package main

import (
	"embed"
	"flag"
	"fmt"
	"github.com/webview/webview"
)

var (
	debug bool = false
)

func init(){
	flag.BoolVar(&debug, "debug", false, "Debug flag")
	flag.Parse()
}

func main(){
	win := webview.New(debug)
	defer win.Destroy()
	win.SetTitle(fmt.Sprintf("Oh my mailbox v%s", VERSION))
	win.SetSize(950, 600, webview.HintNone)
	win.Run()
}
