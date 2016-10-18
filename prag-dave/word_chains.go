package code_kata

import (
	"bufio"
	"container/list"
	"log"
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
	if iter == nil {
		return nil
	}
	v := iter.Value
	q.data.Remove(iter)
	return v
}

type WordGraph struct {
	wordList *list.List
}

func NewWordGraph() *WordGraph {
	return &WordGraph{}
}

func (graph *WordGraph) findShortestPath(fromWord string, toWord string) *list.List {
	var previous map[string]string
	previous = make(map[string]string)

	q := NewQueue()
	q.Enqueue(fromWord)

	for word := q.Dequeue(); q != nil && word != toWord; word = q.Dequeue() {
		innerWord := word.(string)
		nextWords := graph.getNextWords(innerWord)
		for nextWord := nextWords.Front(); nextWord != nil; nextWord = nextWord.Next() {
			val := nextWord.Value.(string)
			if _, ok := previous[val]; ok {
				previous[val] = innerWord
				log.Print(innerWord)
				q.data.PushBack(val)
			}
		}
	}
	return graph.getPath(previous, fromWord, toWord)
}

func (graph *WordGraph) getNextWords(word string) *list.List {
	nextWords := list.New()
	for currWord := graph.wordList.Front(); currWord != nil; currWord = currWord.Next() {
		if graph.getWordDiff(word, currWord.Value.(string)) == 1 {
			nextWords.PushBack(currWord.Value.(string))
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
	log.Print("#####", previous[toWord])
	if _, ok := previous[toWord]; !ok {
		log.Print("---")
		return nil
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
