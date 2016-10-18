package code_kata

import "testing"

func TestCheckWorlds(t *testing.T) {
	s := checkWords("tmp/wordlist.txt")
	t.Errorf("test result: ", s)
}
