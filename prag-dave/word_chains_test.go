package code_kata

import "testing"

func TestFindShortestPath(t *testing.T) {
	graph := new(WordGraph)
	graph.wordList = loadWords("tmp/wordlist.txt", 6)
	val := graph.findShortestPath("cat", "dog")
	for l := val.Front(); l != nil; l = l.Next() {
		t.Errorf(l.Value.(string))
	}
}
