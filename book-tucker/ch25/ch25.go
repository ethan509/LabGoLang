package main

import (
	"fmt"
	"sync"
	"time"
)

func square(wg *sync.WaitGroup, ch chan int) {
	fmt.Println("Waiting...")
	for n := range ch {
		fmt.Println()
		fmt.Printf("Square: %d\n", n*n)
		time.Sleep(time.Second)

		fmt.Println("Waiting in for...")
	}
	fmt.Println("Waiting out of for...")
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}

	close(ch)

	wg.Wait()
}
