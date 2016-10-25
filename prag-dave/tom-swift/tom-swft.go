package code_kata

import (
	"bufio"
	"os"
)

func loadWords(file_path string) {
	dat, err := os.Open(file_path)
	check(err)

	scanner := bufio.newScanner(file_path)

	for scanner.Scan() {
		text := scanner.Text()
	}
}
