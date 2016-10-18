package code_kata

import "testing"

func TestAnagramDict(t *testing.T) {
	str := anagramsDict("tmp/wordlist.txt")
	for _, value := range str {
		if len(value.String()) == 0 {
			t.Errorf("test failed: ", value.String())
		}
	}
}
