package aoc

import "testing"

func TestCalcesult(t *testing.T) {
	result := calcResult("day1.txt")
	if result != 10 {
		t.Errorf("Falied expected %d, got %d", 10, result)
	}
}

func TestCalcesult2(t *testing.T) {
	result := calcResult2("day1_2.txt")
	if result != 10 {
		t.Errorf("Falied expected %d, got %d", 10, result)
	}
}
