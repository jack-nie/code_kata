package aoc

import (
	"bufio"
	"log"
	"os"
)

func loadData(filePath string) []string {
	dat, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	var container []string
	scanner := bufio.NewScanner(dat)
	for scanner.Scan() {
		text := scanner.Text()
		container = append(container, text)
	}
	return container
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	return max(-a, -b)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
