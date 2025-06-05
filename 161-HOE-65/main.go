package main

import "fmt"

func main() {
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", subtract)
	fmt.Printf("%T\n", doMath)

	x := doMath(42, 16, add)
	fmt.Println(x)

	y := doMath(42, 16, subtract)
	fmt.Println(y)
}

func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}


// put this code into it's own file
// Test files live in the same package as the code they test. 
// They are named with the following convention: 
// `filename_test.go` where filename is the name 
// of the file that contains the code you want to test.

/*
package main

import "testing"

func TestAdd(t *testing.T) {
	got := add(7, 5)
	want := 12
	if got != want {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", got, want)
	}
}

func TestSubtract(t *testing.T) {
	got := subtract(7, 5)
	want := 2
	if got != want {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", got, want)
	}
}

func TestDoMath(t *testing.T) {
	got := doMath(7, 5, add)
	want := 12
	if got != want {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", got, want)
	}
}


*/

// to run your test
// at the terminal, aka bash, aka shell
// go test