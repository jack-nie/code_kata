package code_kata

import "testing"

func TestFindShortestPath(t *testing.T) {
	graph = WordGraph(loadWords("tmp/wordlist.txt", 6))
	val := findShortestPath()
	t.Errorf(val)
}
