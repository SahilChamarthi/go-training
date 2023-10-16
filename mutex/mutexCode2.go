package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	var counter = 0

	wg := new(sync.WaitGroup)
	m := sync.Mutex{}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			m.Lock()
			j := counter
			time.Sleep(time.Millisecond)
			j = j + 1

			counter = j
			m.Unlock()
		}()
	}
	wg.Wait()
	log.Println("counter:", counter)
}
