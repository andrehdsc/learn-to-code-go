package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var counter int64
	gs := 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for range gs {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println(atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("End Value:", counter)
}
