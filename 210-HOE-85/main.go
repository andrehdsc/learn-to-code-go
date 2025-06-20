package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	counter := 0
	gs := 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for range gs {
		go func() {
			c := counter
			runtime.Gosched()
			c++
			fmt.Println(c)
			counter = c
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("End Value:", counter)
}
