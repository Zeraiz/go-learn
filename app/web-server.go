package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var (
	mu    sync.Mutex
	count int
)

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Counter: %d\n", count)
	mu.Unlock()
}

func counterWithEchoPart(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "Requested url: %q\n", r.URL.Path)
}

func main() {
	http.HandleFunc("/", counterWithEchoPart)
	http.HandleFunc("/show-counter", handler)

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
