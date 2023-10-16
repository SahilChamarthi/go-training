package main

import (
	"fmt"
	"sync"
)

func main() {

	ch := make(chan int)
	wg := sync.WaitGroup{}
	wgg := sync.WaitGroup{}

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			wgg.Add(1)
			go func(id int) {
				defer wg.Done()
				ch <- id
			}(i)
		}
		wgg.Wait()
		close(ch)
	}()

	go func() {
		defer wg.Done()
		for val := range ch {
			fmt.Println("receiving:", val)
		}
	}()

	wg.Wait()

	fmt.Println("main ends")

}
