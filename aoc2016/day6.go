// --- Day 6: Signals and Noise ---
//
// Something is jamming your communications with Santa. Fortunately, your signal is only partially jammed, and protocol in situations like this is to switch to a simple repetition code to get the message through.
//
// In this model, the same message is sent repeatedly. You've recorded the repeating message signal (your puzzle input), but the data seems quite corrupted - almost too badly to recover. Almost.
//
// All you need to do is figure out which character is most frequent for each position. For example, suppose you had recorded the following messages:
//
// eedadn
// drvtee
// eandsr
// raavrd
// atevrs
// tsrnev
// sdttsa
// rasrtv
// nssdts
// ntnada
// svetve
// tesnvt
// vntsnd
// vrdear
// dvrsen
// enarar
// The most common character in the first column is e; in the second, a; in the third, s, and so on. Combining these characters returns the error-corrected message, easter.
//
// Given the recording in your puzzle input, what is the error-corrected version of the message being sent?
//
// --- Part Two ---
//
// Of course, that would be the message - if you hadn't agreed to use a modified repetition code instead.
//
// In this modified code, the sender instead transmits what looks like random data, but for each character, the character they actually want to send is slightly less likely than the others. Even after signal-jamming noise, you can look at the letter distributions in each column and choose the least common letter to reconstruct the original message.
//
// In the above example, the least common character in the first column is a; in the second, d, and so on. Repeating this process for the remaining characters produces the original message, advent.
//
// Given the recording in your puzzle input and this new decoding methodology, what is the original message that Santa is trying to send?
//
// Your puzzle answer was odqnikqv.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	input := loadData("day6.txt")
	length := len(input)
	var str string
	for i := 0; i < 8; i++ {
		count := make(map[string]int)
		for j := 0; j < length; j++ {
			byt := []byte(input[j])
			count[string(byt[i])] = count[string(byt[i])] + 1
		}
		pl := rankByWordCount(count)
		str = str + pl[0].Key
	}

	fmt.Println(str)
}

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

//Pair for sort
type Pair struct {
	Key   string
	Value int
}

//PairList for sort
type PairList []Pair

func (p PairList) Len() int {
	return len(p)
}

func (p PairList) Less(i, j int) bool {
	return p[i].Value < p[j].Value
}

func (p PairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func rankByWordCount(wordFrequencies map[string]int) PairList {
	pl := make(PairList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = Pair{k, v}
		i++
	}

	sort.Sort(sort.Reverse(pl))
	return pl
}

// part two: in function Less, change from "<" to ">"
