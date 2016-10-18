package code_kata

import "testing"

func TestAnagramDict(t *testing.T) {
	str := anagramsDict("tmp/wordlist.txt")
	for _, value := range str {
		t.Errorf(value.String())
	}
}
