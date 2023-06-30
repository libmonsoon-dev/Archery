package physics_test

import (
	"archery/physics"
	"testing"
)

func TestGetArrowPosition(t *testing.T) {
	got := physics.GetArrowPosition(3, 2, 60, physics.Point{})
	if int(got.X*10_000)/10_000 != 3 {
		t.Fatalf("unexpeced result %+v", got)
	}
}
