package aoc

import "testing"

func TestCalcResult(t *testing.T) {
	result := calcResult("day1.txt")
	if result != 592 {
		t.Errorf("Falied expected %d, got %d", 592, result)
	}
}

func TestCalcResult2(t *testing.T) {
	result := calcResult2("day1_2.txt")
	if result != 241 {
		t.Errorf("Falied expected %d, got %d", 241, result)
	}
}
