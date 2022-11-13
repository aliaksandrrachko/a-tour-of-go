package main

import (
	"fmt"
	"time"
)

func concurrency() {
	// Goroutines - is a lightweight thread managed by the Go runtime
	go say("world")
	say("hello")

	// Channels
	intSliceToSum := []int{
		7, 2, 8, -9, 4, 0, -280, 3,
		-200, 615, -165, 48, 6, 15,
		-5, 6, 5, 54, 6, 1, 32, 2,
	}

	sumChannel := make(chan int)
	go sum(intSliceToSum[:len(intSliceToSum)/2], sumChannel)
	go sum(intSliceToSum[len(intSliceToSum)/2:], sumChannel)
	x, y := <-sumChannel, <-sumChannel // receive from c
	fmt.Println(x, y, x+y)

	// Buffered Channels
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Microsecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}
