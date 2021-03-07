package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func checkError(err error) {
	if err != nil {
		fmt.Printf("Error in fetch: %v", err)
		os.Exit(1)
	}
}

func forceHttps(url string) string {
	res := url
	if strings.HasPrefix(url, "http://") {
		res = strings.Replace(url, "http://", "https://", 8)
	} else if !strings.HasPrefix(url, "https://") {
		res = strings.Join([]string{"https://", url}, "")
	}
	return res
}

func main() {
	for _, url := range os.Args[1:2] {
		resp, err := http.Get(forceHttps(url))
		checkError(err)
		fmt.Printf("HTTP Status code: %d\n", resp.StatusCode)
		f, err := os.Create("./data/parsed-url.html")
		checkError(err)
		_, err = io.Copy(f, resp.Body)
		f.Close()
		checkError(err)
	}
}
