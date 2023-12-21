package simulator

import (
	"testing"
)

func TestSimulator(t *testing.T) {
	want := 1
	num := Simulate()
	if num != want {
		t.Fatalf("Simulate - received %v, expected %v", num, want)
	}
}
