package main

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s\n", time.Now().Format(time.DateTime))
		fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
		fmt.Fprintf(w, "Remote address: %s\n", r.RemoteAddr)
		fmt.Fprintf(w, "Host: %s\n", r.Host)
		fmt.Fprint(w, "---- HEADERS ----\n")

		keys := make([]string, 0, len(r.Header))
		for k := range r.Header {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, k := range keys {
			var sb strings.Builder
			sb.WriteString(k)
			sb.WriteString(": ")
			for _, v := range r.Header[k] {
				sb.WriteString(v)
			}
			sb.WriteString("\n")
			fmt.Fprint(w, sb.String())
		}

		log.Printf(
			"%s \"%s %s %s\" %s \"%s\"\n",
			r.RemoteAddr,
			r.Method,
			r.URL,
			r.Proto,
			fmt.Sprintf("%d", http.StatusOK),
			strings.Join(r.Header["User-Agent"], " "))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
