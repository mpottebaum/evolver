package grid

import (
	"testing"
)

func TestCreateGrid(t *testing.T) {
	startQuad := "Q1"

	result := CreateGrid(startQuad)

	quad1 := result.Quadrants["Q1"]
	if quad1[0] != 0 || quad1[1] != 0 || quad1[2] != Width/2 || quad1[3] != Height/2 {
		t.Logf("Quad 1 - %v", quad1)
		t.Fail()
	}
}
