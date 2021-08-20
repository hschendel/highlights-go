package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	ch := make(chan int) // unbuffered channel
	go consumer(ch)

	var wg sync.WaitGroup
	wg.Add(10)

	for k := 1; k <= 10; k++ {
		go producer(k, k, ch, &wg)
	}

	wg.Wait()
	close(ch)
}

func producer(k, n int, ch chan<- int, wg *sync.WaitGroup) {
	for i := 0; i < n; i++ {
		msToSleep := rand.Int31n(1000)
		time.Sleep(time.Millisecond * time.Duration(msToSleep))
		ch <- k
	}
	wg.Done()
}

func consumer(ch <-chan int) {
	// reads until ch is closed
	for k := range ch {
		fmt.Println(k)
	}
}
