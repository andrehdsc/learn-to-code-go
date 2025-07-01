package main

import (
	"fmt"
)

type customErr struct {
	info string
}

func (err customErr) Error() string {
	return fmt.Sprintf("Look at this error: %v\n", err.info)
}

func main() {
	err1 := customErr{
		info: "Need more coffee",
	}

	foo(err1)

}

func foo(e error) {
	fmt.Println("foo ran! -", e)
}
