package mymath

import (
	"fmt"
	"testing"
)

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{1, 4, 6, 8, 100})
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{1, 4, 6, 8, 100}))
	//Output:
	//6
}

func TestCenteredAvg(t *testing.T) {
	type test struct {
		data   []int
		anwser int
	}
	tests := []test{
		{[]int{10, 20, 40, 60, 80}, 40},
		{[]int{2, 4, 6, 8, 12}, 6},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5},
	}

	for _, xs := range tests {
		v := CenteredAvg(xs.data)
		if v != float64(xs.anwser) {
			t.Error("got", v, "want", xs.anwser)
		}
	}
}
