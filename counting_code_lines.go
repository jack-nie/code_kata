package code_kata

import (
	"bufio"
	"os"
)

var insideQuote, insideBlockComment bool

func countCodeLines(file_path string) int {
	var count int
	count = 0

	dat, _ := os.Open(file_path)
	scanner := bufio.NewScanner(dat)

	for scanner.Scan() {
		text := scanner.Text()
		if isValid(text) {
			count += 1
		}
	}
	return count
}

func isValid(line string) bool {
	var valid bool
	valid = false

	for index := 0; index < len(line); index++ {
		byt := line[index]
		switch byt {
		case ' ':
		case '\n':
		case '\t':
		case '\r':
			break
		case '"':
			if !insideBlockComment {
				valid = true
				insideQuote = !insideQuote
			}
			break
		case '/':
			if isLineComment(line, index) {
				if !insideBlockComment {
					return valid
				}
			}

			if isBlockComment(line, index) {
				insideBlockComment = true
				index++
			} else {
				if !insideBlockComment {
					valid = true
				}
			}
			break
		case '*':
			if insideBlockComment {
				if matchedCharAtIndex('/', line, index+1) {
					insideBlockComment = false
					index++
				}
			} else {
				valid = true
			}
			break
		default:
			if !insideBlockComment {
				valid = true
			}
		}
	}
	return valid
}

func isBlockComment(line string, index int) bool {
	if matchedCharAtIndex('*', line, index+1) {
		if !insideQuote {
			insideBlockComment = true
			return true
		}
	}
	return false
}

func isLineComment(line string, index int) bool {
	return matchedCharAtIndex('/', line, index+1)
}

func matchedCharAtIndex(searchTarget byte, line string, index int) bool {
	if index == len(line) {
		return false
	}
	return searchTarget == line[index]
}
