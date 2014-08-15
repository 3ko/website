package main

import (
	"flag"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
)

var router *mux.Router

const (
	logFile = "access_log.txt"
)

func main() {
	flag.Parse()
	defer glog.Flush()

	router = mux.NewRouter()
	http.HandleFunc("/", httpInterceptor)

	fileServer := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	http.Handle("/static/", fileServer)

	router.HandleFunc("/", home.GetHomePage).Methods("GET")

	http.ListenAndServe(":8080", nil)

}

func httpInterceptor(w http.ResponseWriter, req *http.Request) {
	startTime := time.Now()

	router.ServeHTTP(w, req)

	finishTime := time.Now()
	elapsedTime := finishTime.Sub(startTime)
	LogAccess(w, req, elapsedTime)

	switch req.Method {
	case "GET":
		// We may not always want to StatusOK, but for the sake of
		// this example we will
	case "POST":
		// here we might use http.StatusCreated
	}
}

type accessLog struct {
	ip, method, uri, protocol, host string
	elapsedTime                     time.Duration
}

//LogAccess func
func LogAccess(w http.ResponseWriter, req *http.Request, duration time.Duration) {
	clientIP := req.RemoteAddr

	if colon := strings.LastIndex(clientIP, ":"); colon != -1 {
		clientIP = clientIP[:colon]
	}

	record := &accessLog{
		ip:          clientIP,
		method:      req.Method,
		uri:         req.RequestURI,
		protocol:    req.Proto,
		host:        req.Host,
		elapsedTime: duration,
	}

	writeAccessLog(record)
}

func writeAccessLog(record *accessLog) {
	logRecord := "" + record.ip + " " + record.protocol + " " + record.method + ": " + record.uri + ", host: " + record.host + " (load time: " + strconv.FormatFloat(record.elapsedTime.Seconds(), 'f', 5, 64) + " seconds)"
	glog.Infoln(logRecord)
	glog.Flush()
}
