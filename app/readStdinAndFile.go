package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func readLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)

	if v := f.Name(); v != "/dev/stdin" {
		fmt.Println(filepath.Base(v))
	}

	for input.Scan() {
		if input.Text() == "exit" {
			break
		}
		counts[input.Text()]++
	}
}

// uniq counter from stdin
func main() {

	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		readLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(strings.Join([]string{"./data/", arg}, ""))
			if err != nil {
				fmt.Printf("%v\n", err)
				continue
			}
			readLines(f, counts)
			err = f.Close()
			if err != nil {
				fmt.Printf("%v\n", err)
				continue
			}
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
