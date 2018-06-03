package main

import (
	"fmt"
	"log"
	"net/http"
)

func ipHandler(w http.ResponseWriter, r *http.Request) {
	ip := r.Header.Get("x-forwarded-for")
	fmt.Fprintf(w, "%s", ip)
	log.Println(ip)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "")
}
