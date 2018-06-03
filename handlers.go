package main

import (
	"fmt"
	"log"
	"net/http"
)

func ipHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", r.RemoteAddr)
	log.Println(r.headers["x-forwarded-for"])
}
