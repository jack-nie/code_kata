package code_kata

import (
	"bufio"
	"github.com/fatih/set"
	"os"
)

func checkWords(file_path string) []string {
	var arr []string
	var result []string
	var length int
	set := set.New()

	dat, _ := os.Open(file_path)
	scanner := bufio.NewScanner(dat)

	for scanner.Scan() {
		word := scanner.Text()
		length = len(word)
		if length > 6 {
			continue
		} else if length < 6 {
			set.Add(word)
		} else if length == 6 {
			arr = append(arr, word)
		}
	}

	for _, element := range arr {
		for i := 1; i < 6; i++ {
			first := element[0:i]
			last := element[i:]
			if set.Has(first, last) {
				result = append(result, element)
				break
			}
		}
	}
	return result
}
