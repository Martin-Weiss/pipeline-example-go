package main

import (
	"fmt"
	"net/http"
	"log"
)

const webContent = "Hello SUSE EMEA v2 World! - MARTIN WAS HERE :-)"

func main() {
	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, webContent)
}
