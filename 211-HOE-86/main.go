package main

import (
	"fmt"
	"sync"
)

func main() {

	counter := 0
	gs := 50
	var wg sync.WaitGroup
	wg.Add(gs)

	var mtx sync.Mutex

	for range gs {
		go func() {
			mtx.Lock()
			c := counter
			c++
			fmt.Println(c)
			counter = c
			mtx.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("End Value:", counter)
}