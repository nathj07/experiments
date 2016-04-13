package main

import (
	"net/http"
)

func main() {
	// Simple static webserver:
	// log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("/usr/share/doc"))))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))
	http.ListenAndServe(":3000", nil)
}
