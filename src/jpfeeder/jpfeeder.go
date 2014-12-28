package main

import (
	"bthandler"
	"flag"
	"fmt"
	"log"
	"log/syslog"
	"net"
	"net/http"
	"net/http/fcgi"
	"startones"
)

const APP_VERSION = "0.1"

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")

type FastCGIServer struct{}

func (s FastCGIServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}
	
	startparameters := startones.Start(*golog)
	
	pathinfo := req.Header.Get("X-PATHINFO")
	callback := req.Header.Get("X-CALLBACK")
	quant :=req.Header.Get("X-QUANT")
	
		
	bthandler.BTrequestHandler(*golog,resp, req, "fi_FI", "porno", "site", pathinfo,startparameters,callback,quant)
	

}

func main() {
	flag.Parse() // Scan the arguments list

	if *versionFlag {
		fmt.Println("Version:", APP_VERSION)
	}

	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal(err)
	}
	srv := new(FastCGIServer)
	fcgi.Serve(listener, srv)

}
