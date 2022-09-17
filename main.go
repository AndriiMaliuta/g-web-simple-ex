//package g_web_simple_ex
package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "HELLO!")
	})

	http.ListenAndServe(":4000", nil)
}
