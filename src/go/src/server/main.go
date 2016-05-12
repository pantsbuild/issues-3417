package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html"
	"library"
	"net/http"
)

func main() {
	fmt.Printf(library.HelperFunc())
	r := mux.NewRouter()
	r.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(rw, "Hello, %q", html.EscapeString(req.URL.Path))
	})
	// Don't actually use the router or start a server -- just do enough
	// so all variables + packages are being used or else the Go compiler
	// will complain.
}
