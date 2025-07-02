package main

import "testing"

func TestSqrt(t *testing.T) {
	got, _ := Sqrt(4)
	if got != 42 {
		t.Errorf("Sqrt(4) = %v, want 42", got)
	}
}
