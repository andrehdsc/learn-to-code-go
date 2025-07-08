package word

import (
	"fmt"
	"testing"
)

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count("one Two three")
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount("one Two three")
	}
}

func ExampleCount() {
	fmt.Println(Count("one two tree"))
	//Output:
	//3
}

func TestCount(t *testing.T) {
	wc := Count("one two three")
	if wc != 3 {
		t.Errorf("got: %v, want: 3\n", wc)
	}
}

func TestUseCount(t *testing.T) {
	wc := UseCount("one one one")
	if wc["one"] != 3 {
		t.Errorf("got: %v, want: 3\n", wc["one"])
	}
}
