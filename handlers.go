package main

import (
	"fmt"
	"log"
	"net/http"
)

func ipHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", r.Header["x-forwarded-for"])
	log.Println(r.Header["x-forwarded-for"])
}
