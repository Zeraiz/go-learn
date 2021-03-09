package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func CheckError(err error) {
	if err != nil {
		fmt.Printf("Error in fetch: %v", err)
		os.Exit(1)
	}
}

func ForceHttps(url string) string {
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
		resp, err := http.Get(ForceHttps(url))
		CheckError(err)
		fmt.Printf("HTTP Status code: %d\n", resp.StatusCode)
		f, err := os.Create("./data/parsed-url.html")
		CheckError(err)
		_, err = io.Copy(f, resp.Body)
		f.Close()
		CheckError(err)
	}
}
