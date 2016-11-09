package main

import (
	"fmt"
	http2 "github.com/bradleyfalzon/net/http2"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Add("Link", "</static/main.css>; rel=preload;")
		w.Header().Add("Link", "</static/main.js>; rel=preload;")

		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	s := &http.Server{
		Addr:           ":3003",
		Handler:        nil,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	http2.ConfigureServer(s, nil)

	log.Fatal(s.ListenAndServe())

}
