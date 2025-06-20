package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("A")
		wg.Done()
	}()
	go func() {
		fmt.Println("B")
		wg.Done()
	}()

	fmt.Println("Goroutines:", runtime.NumGoroutine())
	wg.Wait()
}
