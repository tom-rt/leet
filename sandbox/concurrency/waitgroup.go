package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func someCalc(wg *sync.WaitGroup, s int) {
	time.Sleep(time.Second * time.Duration(s))
	fmt.Printf("Done after %d seconds\n", s)
	wg.Done()
}

func WgSandbox() {
	var wg sync.WaitGroup

	wg.Add(1)
	go someCalc(&wg, 2)
	wg.Add(1)
	go someCalc(&wg, 4)
	wg.Wait()
}
