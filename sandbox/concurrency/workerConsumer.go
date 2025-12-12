package concurrency

import (
	"fmt"
	"sync"
)

type Result struct {
	ID  int
	Val int
}

func compute(id int) Result {
	return Result{ID: id, Val: id * id}
}

func WorkerConsumer() {
	const X = 10
	var wg sync.WaitGroup

	results := make(chan Result)

	for i := range X {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("Sent")
			results <- compute(i)
		}(i)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		fmt.Println("Result received:", r)
	}
}
