// Finally, you come across an information kiosk with a list of rooms. Of course, the list is encrypted and full of decoy data, but the instructions to decode the list are barely hidden nearby. Better remove the decoy data first.
//
// Each room consists of an encrypted name (lowercase letters separated by dashes) followed by a dash, a sector ID, and a checksum in square brackets.
//
// A room is real (not a decoy) if the checksum is the five most common letters in the encrypted name, in order, with ties broken by alphabetization. For example:
//
// aaaaa-bbb-z-y-x-123[abxyz] is a real room because the most common letters are a (5), b (3), and then a tie between x, y, and z, which are listed alphabetically.
// a-b-c-d-e-f-g-h-987[abcde] is a real room because although the letters are all tied (1 of each), the first five are listed alphabetically.
// not-a-real-room-404[oarel] is a real room.
// totally-real-room-200[decoy] is not.
// Of the real rooms from the list above, the sum of their sector IDs is 1514.
//
// What is the sum of the sector IDs of the real rooms?

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

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

type Pair struct {
	Key   string
	Value int
}

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

func main() {
	arr := loadData("day4.txt")
	var count int
	for i := range arr {
		count += i
	}
	fmt.Println(count)
}

func loadData(filePath string) []int {
	var tempArr []int
	dat, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(dat)
	reForNum := regexp.MustCompile("[0-9]+")
	reForStr := regexp.MustCompile("[a-z]+")
	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, "-")
		str, numStr := strings.Join(arr[0:len(arr)-1], ""), arr[len(arr)-1]

		num, err := strconv.Atoi(string(reForNum.Find([]byte(numStr))))
		if err != nil {
			log.Fatal(err)
		}

		checkSum := string(reForStr.Find([]byte(numStr)))
		count := make(map[string]int)
		length := len(str)
		byt := []byte(str)
		for j := 0; j < length; j++ {
			count[string(byt[j])] = count[string(byt[j])] + 1
		}
		pl := rankByWordCount(count)
		var tmp []string
		for _, k := range pl {
			tmp = append(tmp, k.Key)
		}

		tmpStr := strings.Join(tmp[:5], "")
		if tmpStr == checkSum {
			tempArr = append(tempArr, num)
		}
	}
	return tempArr
}
