package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {

		time.Sleep(1 * time.Second)

		ch1 <- 10

	}()

	go func() {

		time.Sleep(2 * time.Second)
		ch2 <- 20

	}()

	for i := 1; i <= 2; i++ {
		select {

		case x := <-ch1:
			fmt.Println("receiving:", x)

		case x := <-ch2:
			fmt.Println("receiving:", x)

		}
	}

}
