package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func taskChild(n int) {
	fmt.Println("C: ", n)
	time.Sleep(100 * time.Millisecond)
}

func taskParent(n int) {
	fmt.Println("P: ", n)
	time.Sleep(200 * time.Millisecond)
}

func child(ctx context.Context) {
	child, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	done := false
	go func() {
		for i := 0; !done; i++ {
			taskChild(i)
		}
	}()

	<-child.Done()
	fmt.Println("C Done")
	done = true
}

func parent(ctx context.Context, wg *sync.WaitGroup) {
	parent, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	//start child task
	go child(parent)

	done := false
	go func() {
		for i := 0; !done; i++ {
			taskParent(i)
		}
	}()

	<-parent.Done()
	fmt.Println("P Done")
	done = true
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	go parent(ctx, &wg)

	<-ctx.Done()
	fmt.Println("timeout ...")

	cancel()
	fmt.Println("cancel context ...")

	wg.Done()
	fmt.Println("exit now...")
}
