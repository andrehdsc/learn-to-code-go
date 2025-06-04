package main

import (
	"fmt"
)

func main() {
	fmt.Println("first")
	defer fmt.Println("forth")  // this will run last
	defer fmt.Println("third")  // this will run second-to-last
	defer fmt.Println("second") // this will run third-to-last
}

// The DEFER keyword works in a "last in, first out" method.
