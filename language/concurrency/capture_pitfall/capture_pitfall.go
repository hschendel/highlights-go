package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wrong()
	fmt.Println()
	correct()
}

func wrong() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 1; i <= 10; i++ {
		go func() {
			time.Sleep(time.Millisecond * 10)
			fmt.Println(i)
			wg.Done()
		}()
	}
	wg.Wait() // if main() exits, the program exits, so we better wait!
}

func correct() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 1; i <= 10; i++ {
		go func(i int) {
			time.Sleep(time.Millisecond * 10)
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
