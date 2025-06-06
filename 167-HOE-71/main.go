package main

import "fmt"

func main() {
	x := foo(2, 2, adder)
	fmt.Println(x)
}

func adder(x int, y int) int {
	return x + y
}

func foo(x int, y int, f func(int, int) int) int {
	return f(x, y)
}
