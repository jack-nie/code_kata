package code_kata

import (
	"bufio"
	"container/list"
	"os"
)

type Queue struct {
	data *list.List
}

func NewQueue() *Queue {
	q := new(Queue)
	q.data = list.New()
	return q
}

func (q *Queue) Enqueue(v interface{}) {
	q.data.PushBack(v)
}

func (q *Queue) Dequeue() interface{} {
	iter := q.data.Front()
	v := iter.Value
	q.data.Remove(iter)
	return v
}

type WordGraph struct {
	wordList *list.List
}

func (graph *WordGraph) findShortestPath(fromWord string, toWord string) *list.List {
	var previous map[string]string
	var word string
	previous = make(map[string]string)

	q := NewQueue()
	q.Enqueue(fromWord)

	for word := q.Dequeue().(string); word != toWord && q.data.Len() != 0; word = q.Dequeue().(string) {
		nextWords := graph.getNextWords(word)
		for nextWord := nextWords.Front(); nextWord != nil; nextWord = nextWord.Next() {
			val := nextWord.Value.(string)
			if _, ok := previous[val]; ok {
				previous[val] = word
				q.data.PushBack(val)
			}
		}
	}
	return graph.getPath(previous, fromWord, toWord)
}

func (graph *WordGraph) getNextWords(word string) *list.List {
	var nextWords *list.List
	for currWord := graph.wordList.Front(); currWord != nil; currWord = currWord.Next() {
		if graph.getWordDiff(word, currWord.Value.(string)) == 1 {
			nextWords.PushBack(currWord)
		}

	}
	return nextWords
}

func (graph *WordGraph) getWordDiff(word1 string, word2 string) int {
	var diff int
	diff = 0
	for i := 0; i < len(word1); i++ {
		if string(word1[i]) != string(word2[i]) {
			diff++
		}
	}
	return diff
}

func (graph *WordGraph) getPath(previous map[string]string, fromWord string, toWord string) *list.List {
	path := list.New()
	if _, ok := previous[toWord]; ok {
		return path
	}

	word := toWord
	for word != fromWord {
		path.PushFront(word)
		word = previous[word]
	}

	path.PushFront(fromWord)
	return path
}

func loadWords(file_path string, wordLength int) *list.List {
	wordList := list.New()
	dat, err := os.Open(file_path)
	check(err)

	scanner := bufio.NewScanner(dat)

	for scanner.Scan() {
		word := scanner.Text()
		if len(word) == wordLength {
			wordList.PushBack(word)
		}
	}

	return wordList
}
