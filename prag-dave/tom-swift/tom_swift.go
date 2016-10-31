package code_kata

import (
	"bufio"
	"os"
	"strings"
)

func loadWords(file_path string) []string {
	dat, _ := os.Open(file_path)

	scanner := bufio.NewScanner(dat)

	var text []string

	for scanner.Scan() {
		text = strings.Fields(scanner.Text())
	}
	return text
}

func makeDict(file_path string) map[string][]string {
	dict := make(map[string][]string)
	text := loadWords(file_path)

	for i := 0; i < len(text)-2; i++ {
		key := strings.Join(text[i:(i+2)], " ")
		dict[key] = append(dict[key], text[i+2])
	}

	return dict
}
