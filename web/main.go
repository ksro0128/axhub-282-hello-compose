package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	greeting := os.Getenv("GREETING")
	if greeting == "" {
		greeting = "hello"
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s — redis=%s\n", greeting, os.Getenv("REDIS_ADDR"))
	})
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	})
	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
