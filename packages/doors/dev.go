//go:build dev
// +build dev

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

const (
	defaultPort = 8080
)

var (
	port = defaultPort
)

func init() {
	if pStr := os.Getenv("PORT"); pStr != "" {
		p, err := strconv.Atoi(pStr)
		if err != nil {
			log.Fatalf("invalid port: %v\n", pStr)
		}
		port = p
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		qp := r.URL.Query()

		params := make(map[string]interface{}, len(qp))
		for k := range qp {
			params[k] = qp.Get(k)
		}

		w.Header().Set("Content-Type", "text/html")
		payload := Main(params)
		if b, ok := payload["error"]; ok {
			fmt.Fprint(w, b)
			return
		}
		if b, ok := payload["body"]; ok {
			fmt.Fprint(w, b)
		}
	})

	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Running on %s\n", addr)
	http.ListenAndServe(addr, nil)
}
