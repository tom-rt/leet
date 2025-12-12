package concurrency

import (
	"fmt"
	"time"
)

func SelectSandbox() {
	var ch1 chan string = make(chan string)
	var ch2 chan string = make(chan string)

	// send to ch1 after 1 second
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "message from ch1"
	}()

	// send to ch2 after 2 seconds
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "message from ch2"
	}()

	// select chooses the channel that becomes ready first
	// Here only chan1 will receive, value passed to ch2 is lost
	select {
	case msg := <-ch1:
		fmt.Println("Received:", msg)
	case msg := <-ch2:
		fmt.Println("Received:", msg)
	}

	var ch3 chan string = make(chan string)
	var ch4 chan string = make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		ch3 <- "ch3 value"
		fmt.Println("val sent to ch3")
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch4 <- "ch4 value"
		fmt.Println("val sent to ch4")
	}()

	// in this case both values will be received
	ret := <-ch3
	fmt.Println(ret)
	ret = <-ch4
	fmt.Println(ret)

}
