package aoc

import "testing"

func TestGetLenOverLaps(t *testing.T) {
	result1, result2 := getLenOverlaps()
	if result1 != 592 {
		t.Errorf("Falied expected %d, got %d", 592, result1)
	}
	if result2 != 592 {
		t.Errorf("Falied expected %d, got %d", 592, result2)
	}
}
