package code_kata

import (
	"bufio"
	"github.com/fatih/set"
	"log"
	"os"
	"sort"
)

type sortRunes []rune

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func anagramsDict(file_path string) map[string]*set.Set {
	var dict map[string]*set.Set
	dict = make(map[string]*set.Set)

	dat, err := os.Open(file_path)

	if err != nil {
		log.Fatal("File read failed: ", err)
	}

	scanner := bufio.NewScanner(dat)

	for scanner.Scan() {
		orig := scanner.Text()
		str := SortString(orig)
		if dict[str] == nil {
			dict[str] = set.New(orig)
		} else {
			dict[str].Add(orig)
		}
	}

	return dict
}
