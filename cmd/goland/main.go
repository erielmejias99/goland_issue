package main

import (
	"fmt"
	"net/http"
)

func main() {

	r := http.NewServeMux()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s", "World")
	})
	fmt.Print(http.ListenAndServe(":4000", r))
}
