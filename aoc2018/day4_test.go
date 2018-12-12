package aoc

import "testing"

func TestReposeRecord(t *testing.T) {
	s := reposeRecord("day4.txt")
	if s != 99759 {
		t.Errorf("Failed, expected %d, got %d", 99759, s)
	}
}

func TestReposeRecordTwo(t *testing.T) {
	s := reposeRecordTwo("day4_2.txt")
	if s != 97884 {
		t.Errorf("Failed, expected %d, got %d", 97884, s)
	}
}
