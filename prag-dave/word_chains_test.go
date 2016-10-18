package code_kata

import "testing"

func TestFindShortestPath(t *testing.T) {
	graph := new(WordGraph)
	graph.wordList = loadWords("tmp/wordlist.txt", 3)
	val := graph.findShortestPath("cat", "fat")
	if val != nil {
		for l := val.Front(); l != nil; l = l.Next() {
			t.Errorf("------")
		}

	}
}
