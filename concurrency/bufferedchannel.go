package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := &sync.WaitGroup{}
	wg.Add(1)

	ch := make(chan int, 1) // buffered channel

	go func() {
		defer wg.Done()
		ch <- 20 // sender

	}()

	x := <-ch // receiver

	fmt.Println(x)

	wg.Wait()
	fmt.Println("enfd of main")
}
