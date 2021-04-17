package main

import (
	"fmt"
	"tempconv"
)

func main() {
	celsiusZero := tempconv.Celsius(22)
	fmt.Printf("%v\n", celsiusZero)
}
