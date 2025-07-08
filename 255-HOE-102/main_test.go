package main

import (
	"fmt"
	"testing"

	"github.com/andrehdsc/learn-to-code-go/255-HOE-102/word"
)

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		word.Count("one Two three")
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		word.UseCount("one Two three")
	}
}

func ExampleCount() {
	fmt.Println(word.Count("one two tree"))
	//Output:
	//3
}

func TestCount(t *testing.T) {
	wc := word.Count("one two three")
	if wc != 3 {
		t.Errorf("got: %v, want: 3\n", wc)
	}
}

func TestUseCount(t *testing.T) {
	wc := word.UseCount("one one one")
	if wc["one"] != 3 {
		t.Errorf("got: %v, want: 3\n", wc["one"])
	}
}


