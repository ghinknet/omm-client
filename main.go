
package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"net"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/webview/webview"
)


//go:generate bash ./omm-front/build.sh
//go:embed all:omm-front/dist
var _dist embed.FS
var dist = func()(fs.FS){
	f, err := fs.Sub(_dist, "omm-front/dist")
	must(err)
	return f
}()

var (
	debug bool = false
	hostIP string = "127.0.0.1"
)

func init(){
	flag.BoolVar(&debug, "debug", debug, "Debug flag")
	flag.StringVar(&hostIP, "host", hostIP, "Web app host IP")
	flag.Parse()
}

func makeRouter()(r *chi.Mux){
	r = chi.NewRouter()
	r.Handle("/W/*", http.StripPrefix("/W", http.FileServer(http.FS(dist))))
	return
}

func main(){
	listener, err := net.ListenTCP("tcp4", &net.TCPAddr{IP: net.ParseIP(hostIP)})
	must(err)
	addr := listener.Addr().(*net.TCPAddr)
	if debug {
		fmt.Println("Server addr:", addr.String())
	}
	webRouter := makeRouter()
	go func(){
		svr := &http.Server{
			Handler:        webRouter,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
		}
		if err := svr.Serve(listener); err != nil {
			panic(err)
		}
	}()

	win := webview.New(debug)
	defer win.Destroy()
	win.SetTitle(fmt.Sprintf("Oh my mailbox v%s", VERSION))
	win.SetSize(950, 600, webview.HintNone)
	navuri := "http://" + addr.String() + "/W/"
	if debug {
		fmt.Println("Navigate url:", navuri)
	}
	win.Navigate(navuri)
	win.Run()
}

func must(err error){
	if err != nil {
		panic(err)
	}
}
