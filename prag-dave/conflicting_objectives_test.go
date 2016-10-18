package code_kata

import "testing"

func TestCheckWorlds(t *testing.T) {
	s := checkWords("tmp/wordlist.txt")
	if len(s) == 0 {
		t.Errorf("test failed: ", s)
	}
}
