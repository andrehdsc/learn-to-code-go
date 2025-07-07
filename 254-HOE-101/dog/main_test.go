package dog

import (
	"fmt"
	"testing"

	//"github.com/andrehdsc/learn-to-code-go/254-HOE-101/dog"
)

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(2)
	}
}

func ExampleYears() {
	fmt.Println(YearsTwo(2))
	//Output:
	//14	
}

func TestYears(t *testing.T) {
	v := Years(2)
	if v != 14 {
		t.Errorf("Test failed! Expected: 14, got: %v\n", v)
	}
}
func TestYearsTwo(t *testing.T) {
	v := YearsTwo(2)
	if v != 14 {
		t.Errorf("Test failed! Expected: 14, got: %v\n", v)
	}
}