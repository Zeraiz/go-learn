package main

import (
	"fmt"
	"os"
)

func main() {
	var s string
	sep := " "
	if len(os.Args) > 1 {
		s += os.Args[1]
		for i := 2; i < len(os.Args); i++ {
			s += sep + os.Args[i]
		}
	}

	fmt.Println(s)
}
