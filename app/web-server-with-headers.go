package main

import (
	"fmt"
	"log"
	"net/http"
	"sort"
)

func handler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Method: %q, Url: %q, Proto %q\n", r.Method, r.URL, r.Proto)

	headers := make([]string, 0, len(r.Header))
	for k, _ := range r.Header {
		headers = append(headers, k)
	}

	sort.Strings(headers)

	for _, v := range headers {
		_, _ = fmt.Fprintf(w, "Header[%q] = %q\n", v, r.Header[v])
	}
	fmt.Fprintf(w, "Host: %q\n", r.Host)
	fmt.Fprintf(w, "Remote addrs: %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form field [%q] = %q\n", k, v)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8081", nil))
}
