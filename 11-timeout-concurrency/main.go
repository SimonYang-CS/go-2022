package main

import (
	"fmt"
	"time"
)

func process(n int) {
	for i := 0; i < n; i++ {
		fmt.Println("i: ", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func timeout(n int) {
	done := make(chan struct{})

	go func() {
		process(n)
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("done:")
		return
	case <-time.After(2 * time.Second):
		fmt.Println("timeout")
		return
	}
}

func main() {
	timeout(10)
	timeout(100)
}
