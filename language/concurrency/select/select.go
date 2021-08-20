package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	ch := make(chan int) // unbuffered channel
	factorCh := make(chan int)

	go consumer(ch, factorCh)

	var wg sync.WaitGroup
	wg.Add(10)

	for k := 1; k <= 10; k++ {
		go producer(k, k, ch, &wg)
	}

	time.Sleep(time.Millisecond * 400)
	factorCh <- 100

	time.Sleep(time.Millisecond * 400)
	factorCh <- -1

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

func consumer(ch, factorCh <-chan int) {
	factor := 1
	for {
		select {
		case k, ok := <-ch:
			if !ok {
				// ch is closed
				return
			}
			fmt.Println(factor * k)
		case newFactor := <-factorCh:
			factor = newFactor
		}
	}
}
