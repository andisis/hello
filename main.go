package main

import (
	"fmt"
	"log"
	"net/http"
)

const resp = "Halo, nama saya %s"

func main() {
	http.HandleFunc("/", greetingHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func greetingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, resp, r.URL.Path[1:])
}
