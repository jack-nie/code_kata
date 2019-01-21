package main

import (
	"bufio"
	"log"
	"os"
)

type stringInterner map[string]string

func (si stringInterner) InternBytes(b []byte) string {
	if interned, ok := si[string(b)]; ok {
		return interned
	}
	s := string(b)
	si[s] = s
	return s
}

func main() {
	f, err := os.Open("1984.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	si := stringInterner{}
	var words []string
	for scanner.Scan() {
		words = append(words, si.InternBytes(scanner.Bytes())) // intern words
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Print(len(words))
}
