package main

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func printNameFunc(i interface{}) {
	fmt.Println(runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name())
}

func checkTimeForFunction(f func()) {
	start := time.Now()
	f()
	end := time.Now()
	fmt.Println(strings.Join([]string{"Time spended: ", fmt.Sprintf("%f", end.Sub(start).Seconds()), "\n"}, ""))
}

func printArgsOneByOneLine() {
	for i, el := range os.Args[1:] {
		fmt.Println("Index: " + strconv.Itoa(i) + " Argument: " + el)
	}
}

func printArgsOneByOneLineASJoin() {
	for i, el := range os.Args[1:] {
		fmt.Println(strings.Join([]string{"Index: ", strconv.Itoa(i), " Argument: ", el}, ""))
	}
}

func main() {
	checkTimeForFunction(func() {
		printNameFunc(printNameFunc)
	})
	checkTimeForFunction(func() {
		printArgsOneByOneLine()
	})
	checkTimeForFunction(func() {
		printArgsOneByOneLineASJoin()
	})
}
