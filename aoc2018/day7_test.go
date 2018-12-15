package aoc

import (
	"fmt"
	"testing"
)

func TestInstructionOrder(t *testing.T) {
	result := instructionOrder("day7.txt")
	s := fmt.Sprintf("%s", result)
	if s != "sss" {
		t.Errorf("Failed, expected %s, got %s", "sss", s)
	}
}
