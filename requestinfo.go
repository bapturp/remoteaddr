package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/", RequestInfo)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func RequestInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\n", time.Now().Format(time.DateTime))
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL.Path, r.Proto)
	fmt.Fprintf(w, "%-15s %s\n", "Remote address:", r.RemoteAddr)
	fmt.Fprintf(w, "%-15s %s\n", "Host:", r.Host)

	if len(r.Header) > 0 {
		fmt.Fprint(w, "Headers:\n")
		for key, values := range r.Header {
			fmt.Fprintf(w, "  %s: %s\n", key, strings.Join(values, " "))
		}
	}

	log.Printf(
		"%s \"%s %s %s\" %d \"%s\"\n",
		r.RemoteAddr,
		r.Method,
		r.URL,
		r.Proto,
		http.StatusOK,
		strings.Join(r.Header["User-Agent"], " "))
}
