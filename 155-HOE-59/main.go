package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}

	fmt.Println("foo: ", foo(x...))
	fmt.Println("bar: ", bar(x))
}

func foo(x ...int) int {
	total := 0
	for _, v := range x {
		total += v	
	}
	return total
}

func bar(x []int) int {
	total := 0
	for _, v := range x {
		total += v	
	}
	return total
}