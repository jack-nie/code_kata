package aoc

import "testing"

func TestGetLenOverlaps(t *testing.T) {
	result := getLenOverlaps()
	if result != 124850 {
		t.Errorf("Falied expected %d, got %d", 124850, result)
	}
}

func TestGetLenOverlapsTwo(t *testing.T) {
	result := getLenOverlapsTwo()
	if result != 1097 {
		t.Errorf("Falied expected %d, got %d", 1097, result)
	}
}
