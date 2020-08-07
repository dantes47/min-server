package main

import (
	"fmt"
	"net/http"
)

func server(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello there! I like %s.", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", server)
	http.ListenAndServe(":4000", nil)
}
