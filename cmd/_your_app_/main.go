package main

import (
	"fmt"
	"net/http"
)

func dummy(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "dummy\n")
}

func main() {
	http.HandleFunc("/", dummy)
	http.ListenAndServe(":8090", nil)
}
