package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"io/ioutil"
)

func ipHtmlHandler(w http.ResponseWriter, r *http.Request) {
	html := getHtml(r.Header.Get("x-forwarded-for"))
	fmt.Fprintf(w, "%s", html)
}

func ipHandler(w http.ResponseWriter, r *http.Request) {
	ip := r.Header.Get("x-forwarded-for")
	fmt.Fprintf(w, "%s", ip)
	log.Println(ip)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "")
}

func getHtml(content string) string {
	b, err := ioutil.ReadFile("html/layout.html")
	if err != nil {
		fmt.Println(err)
	}

	html := string(b)

	html = strings.Replace(html, "{{ content }}", content, 1)

	return html
}
