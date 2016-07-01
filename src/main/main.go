package main

import (
	"log"
	"net/http"
)

// someFunc Handler function for "/" pattern
func someFunc(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello universe! Using NewServerMux"))
}

func main() {
	serverMux := http.NewServeMux()

	serverMux.HandleFunc("/", someFunc)
	err := http.ListenAndServe(":8080", serverMux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
