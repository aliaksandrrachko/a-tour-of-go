package main

import (
	"fmt"
	"sync"
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

	// sync.Mutex
	safeCounter := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go safeCounter.Inc("someKey")
	}

	time.Sleep(time.Second)
	fmt.Println(safeCounter.Value("someKey"))

	// Exercise: Web Crawler
	Crawl("https://golang.org/", 4, fetcher)

	fmt.Println("Finish")
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

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map.c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	cash := ConcurrentMap{v: make(map[string]FetchResult)}
	var wg sync.WaitGroup

	RecursiveCrawl(url, depth, fetcher, &cash, &wg)
	wg.Wait()
}

func RecursiveCrawl(url string, depth int,
	fetcher Fetcher,
	cash *ConcurrentMap,
	wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	if depth <= 0 {
		return
	}
	if cash.ContainsKey(url) {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	fetchResult := FetchResult{body, urls, err}
	if fetchResult.err != nil {
		fmt.Println(fetchResult.err)
		return
	}
	cash.Put(url, fetchResult)
	fmt.Printf("found: %s %q\n", url, fetchResult.body)
	for _, u := range fetchResult.urls {
		RecursiveCrawl(u, depth-1, fetcher, cash, wg)
	}
}

// ConcurrentMap
type ConcurrentMap struct {
	mu sync.Mutex
	v  map[string]FetchResult
}

func (c *ConcurrentMap) Put(key string, value FetchResult) {
	c.mu.Lock()
	c.v[key] = value
	c.mu.Unlock()
}

func (c *ConcurrentMap) ContainsKey(key string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.v[key]; ok {
		return true
	}
	return false
}

type FetchResult struct {
	body string
	urls []string
	err  error
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
