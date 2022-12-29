package main

import (
	"fmt"
	"sync"
)

func write(wg *sync.WaitGroup, ch chan int) {
	ch <- 100
	wg.Done()
}

func read(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	fmt.Println(<-ch)
}

func main() {
	ch := make(chan int)

	wg := &sync.WaitGroup{}

	wg.Add(2)

	go write(wg, ch)

	go read(wg, ch)

	wg.Wait()
}
