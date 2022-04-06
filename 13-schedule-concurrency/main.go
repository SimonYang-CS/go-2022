package main

import "time"

func task1() {
	time.Sleep(5 * time.Second)
}

func task2() {
	time.Sleep(2 * time.Second)
}

func task3() {
	time.Sleep(2 * time.Second)
}

func main() {
	for i := 0; i < 10; i++ {
		go task1()
		go task2()
		go task3()

		time.Sleep(time.Second)
	}
}
