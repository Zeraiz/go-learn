package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
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

func fetchAsync(url string, ch chan<- string) {
	startTime := time.Now()
	resp, err := http.Get(forceHttps(url))
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	urlRg := regexp.MustCompile("[^a-zA-Z0-9.\\-]")
	escapedFileName := urlRg.ReplaceAllString(url, "_")
	file, err := os.Create(fmt.Sprintf("./data/%s.html", escapedFileName))
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(file, resp.Body)
	_ = resp.Body.Close()
	_ = file.Close()
	if err != nil {
		ch <- fmt.Sprintf("Error while copy url: %s, error: %v")
		return
	}
	ch <- fmt.Sprintf("%.2fs %7d %s", time.Since(startTime).Seconds(), nbytes, url)
}

func main() {
	startTime := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetchAsync(url, ch)
	}
	for range os.Args[1:] {
		println(<-ch)
	}
	endTime := time.Since(startTime).Seconds()
	fmt.Printf("%.2fs elapsed\n", endTime)
}
