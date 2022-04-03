package main

import (
	"fmt"
	"sync"
	"time"
)

func step1(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Step1 Sleep 1s")
	time.Sleep(time.Second)
	fmt.Println("Step1 done")
}

func step2(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Step2 Sleep 3s")
	time.Sleep(3 * time.Second)
	fmt.Println("step2 done")
}

func step3(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Step3 Sleep 2s")
	time.Sleep(2 * time.Second)
	fmt.Println("step3 done")
}

func main() {
	var wg sync.WaitGroup

	wg.Add(3)

	go step1(&wg)
	go step2(&wg)
	go step3(&wg)

	wg.Wait()
	fmt.Println("done")
}
