package main

import (
	"fmt"
	"time"
)

func main() {
	timeFunc(foo)
}

func foo() {
	fmt.Println("Foo started...")
	time.Sleep(2 * time.Second)
	fmt.Println("Foo completed!")
}

func timeFunc(fn func()) {
	start := time.Now()
	fn()
	elapsed := time.Since(start)
	fmt.Println("The elapsed time to run the function was: ", elapsed)
}
