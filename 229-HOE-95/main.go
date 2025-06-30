package main

import "fmt"

func main() {
	ch := make(chan int)

	for range 10 {
		go adder(ch)
	}

	for i := range 100 {
		fmt.Println(i, <- ch)
	}
	close(ch)
}

func adder(c chan<- int) {
	for i := range 10 {
		c <- i
	}
}
