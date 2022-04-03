package main

import (
	"fmt"
	"runtime/debug"
)

func div(x int, y int) {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()

	fmt.Println(x / y)
}

func main() {
	for _, val := range []int{10, 20, 0, 40} {
		div(200, val)
	}
}
