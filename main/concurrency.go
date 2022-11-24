package main

import (
	"fmt"
	"time"

	"golang.org/x/tour/tree"
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
	bufferedChannelExample := make(chan int, 2)
	bufferedChannelExample <- 1
	bufferedChannelExample <- 2
	fmt.Println(<-bufferedChannelExample)
	fmt.Println(<-bufferedChannelExample)

	// Range and Close
	fibonacciChan := make(chan int, 10)
	go fibonacciWithChan(cap(fibonacciChan), fibonacciChan)
	for i := range fibonacciChan {
		fmt.Println(i)
	}

	// Select
	fibonacciChanForSelectExample := make(chan int)
	quitChan := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-fibonacciChanForSelectExample)
		}
		quitChan <- 0
	}()
	fibonacciWithChanAndSelect(fibonacciChanForSelectExample, quitChan)

	// Default Selection
	DefaultSelection()

	// Exercise: Equivalent Binary Trees
	channelForRandomBinaryTree := make(chan int)
	go Walk(tree.New(1), channelForRandomBinaryTree)
	for i := range channelForRandomBinaryTree {
		fmt.Println(i)
	}

	fmt.Println(
		Same(tree.New(1), tree.New(1)),
	)
	fmt.Println(
		Same(tree.New(1), tree.New(2)),
	)
}

func DefaultSelection() bool {
	tickDefaultSelection := time.Tick(100 * time.Millisecond)
	boomDefaultSelection := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tickDefaultSelection:
			fmt.Println("tick.")
		case <-boomDefaultSelection:
			fmt.Println("BOOM!")
			return true
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
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

func fibonacciWithChan(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func fibonacciWithChanAndSelect(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

// Walk walks the three t sending all values
// from the three to the change ch.
func Walk(t *tree.Tree, ch chan int) {
	WalkRecursive(t, ch)
	close(ch)
}

func WalkRecursive(t *tree.Tree, ch chan int) {
	if t != nil {
		WalkRecursive(t.Left, ch)
		ch <- t.Value
		WalkRecursive(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		n1, ok1 := <-ch1
		n2, ok2 := <-ch2
		if ok1 != ok2 || n1 != n2 {
			return false
		}
		if !ok1 {
			break
		}
	}
	return true
}
