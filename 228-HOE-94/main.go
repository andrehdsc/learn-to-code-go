package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() {
		for i := range 100 {
			ch <- i
		}
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}
}
