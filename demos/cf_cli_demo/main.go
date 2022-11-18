package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var name string

	found := r.URL.Query().Get("name")
	if found != "" {
		name = found
	} else {
		name = "world"
	}

	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
	var port = os.Getenv("PORT")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}
