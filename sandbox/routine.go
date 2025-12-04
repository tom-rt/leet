package main

import (
	"fmt"
	"time"
)

func calcA(c chan int) {
	time.Sleep(time.Second * 5)
	c <- 30
}

func calcB(c chan int) {
	time.Sleep(time.Second * 2)
	c <- 70
}

func simpleRoutine() {
	c := make(chan int)
	go calcA(c)
	go calcB(c)
	x, y := <-c, <-c
	fmt.Println(x, y)
}
