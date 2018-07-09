package main

import (
	"io"
	"log"
	"net/http"
)

func server(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello")
}

func main() {
	http.HandleFunc("/", server)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
