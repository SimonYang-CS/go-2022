package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var count int32

const LEN = 100

func newFunc() interface{} {
	return atomic.AddInt32(&count, 1)
}

func task(pool *sync.Pool, wg *sync.WaitGroup) {
	v := pool.Get()

	fmt.Printf("v: %v\n", v)
	pool.Put(v)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(LEN)

	pool := sync.Pool{
		New: newFunc,
	}

	for i := 0; i < LEN; i++ {
		go task(&pool, &wg)
	}

	wg.Wait()
}
