package aoc

import "testing"

func TestCalcBoxIdCheckSum(t *testing.T) {
	result := calcBoxIdCheckSum("day2.txt")
	if result != 5166 {
		t.Errorf("Failed, expected %d, got %d", 5166, result)
	}
}

func TestCalcBoxIdCheckSumTwo(t *testing.T) {
	result := calcBoxIdCheckSumTwo("day2_2.txt")
	if result != "cypueihajytordkgzxfqplbwn" {
		t.Errorf("Failed, expected %s, got %s", "cypueihajytordkgzxfqplbwn", result)
	}
}
