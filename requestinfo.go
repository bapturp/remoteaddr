package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

type RequestInfo struct {
	Time       time.Time
	Method     string
	Path       string
	Proto      string
	RemoteAddr string
	Host       string
	Headers    map[string]string
}

func (ri RequestInfo) DefaultResponse() string {
	var s string
	s += fmt.Sprintf("%s\n", ri.Time.Format(time.DateTime))
	s += fmt.Sprintf("%s %s %s\n", ri.Method, ri.Path, ri.Proto)
	s += fmt.Sprintf("%-15s %s\n", "Remote address:", ri.RemoteAddr)
	s += fmt.Sprintf("%-15s %s\n", "Host:", ri.Host)
	s += "Headers:"
	for k, v := range ri.Headers {
		s += fmt.Sprintf("  %s: %s\n", k, v)
	}
	return s
}

func (ri RequestInfo) JSONResponse() string {
	s, err := json.Marshal(ri)
	if err != nil {
		log.Printf("Error while marshalling: %v", err)
		return ""
	}
	return string(s)
}

var addr string = ":8080"

func main() {
	http.HandleFunc("/", HandlerRequestInfo)
	log.Printf("Server listening on %v\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func HandlerRequestInfo(w http.ResponseWriter, r *http.Request) {
	ri := RequestInfo{
		Time:       time.Now(),
		Method:     r.Method,
		Path:       r.URL.Path,
		Proto:      r.Proto,
		RemoteAddr: r.RemoteAddr,
		Host:       r.Host,
		Headers:    map[string]string{},
	}

	for key, values := range r.Header {
		ri.Headers[key] = strings.Join(values, " ")
	}

	if responseQuery := r.URL.Query().Get("r"); responseQuery == "json" {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, ri.JSONResponse())
	} else {
		fmt.Fprint(w, ri.DefaultResponse())
	}

	logAccess(ri)
}

func logAccess(ri RequestInfo) {
	log.Printf(
		"%s \"%s %s %s\" %d \"%s\"\n",
		ri.RemoteAddr,
		ri.Method,
		ri.Path,
		ri.Proto,
		http.StatusOK,
		ri.Headers["User-Agent"])
}
