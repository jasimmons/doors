//go:build dev
// +build dev

package main

import (
	"fmt"
	"net/http"
)

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

	fmt.Println("Running on port :8080")
	http.ListenAndServe(":8080", nil)
}
