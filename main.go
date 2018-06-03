package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	port = "8080"
)

func main() {
	fmt.Println("Starting...")
	http.HandleFunc("/ip", ipHandler)
	http.HandleFunc("/ip/html", ipHtmlHandler)
	http.HandleFunc("/", homeHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
