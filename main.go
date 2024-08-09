package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", serveRoot)

	err := http.ListenAndServe(":3000", http.DefaultServeMux)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func serveRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
