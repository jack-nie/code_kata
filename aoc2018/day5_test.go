package aoc

import "testing"

func TestReact(t *testing.T) {
	result := react("day5.txt")
	if result != 9238 {
		t.Errorf("Failed, expect %d, got %d", 9238, result)
	}
}
