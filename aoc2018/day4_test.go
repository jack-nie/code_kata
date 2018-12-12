package aoc

import "testing"

func TestReposeRecord(t *testing.T) {
	s := reposeRecord("day4.txt")
	if s != 12 {
		t.Errorf("Failed, expected %d, got %d", 12, s)
	}
}
