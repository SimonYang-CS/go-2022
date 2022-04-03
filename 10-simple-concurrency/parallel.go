package main

import (
	"fmt"
	"time"
)

const LEN = 10

func process(v int) int {
	time.Sleep(10 * time.Millisecond)
	return v * v
}

func parallel(in <-chan int, out chan<- int) {
	go func() {
		for input := range in {
			out <- process(input)
		}

		//all goroutines are asleep
		close(out)

		//can't close receive-only channel
	}()
}

func TryParallel() {
	fmt.Println("TryParallel")

	input := make(chan int, LEN)
	output := make(chan int)

	for i := 0; i < LEN; i++ {
		input <- i
	}
	close(input)

	parallel(input, output)

	for value := range output {
		fmt.Println(value)
	}
}
