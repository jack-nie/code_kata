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
// --- Part Two ---
//
// With all the decoy data out of the way, it's time to decrypt this list and get moving.
//
// The room names are encrypted by a state-of-the-art shift cipher, which is nearly unbreakable without the right software. However, the information kiosk designers at Easter Bunny HQ were not expecting to deal with a master cryptographer like yourself.
//
// To decrypt a room name, rotate each letter forward through the alphabet a number of times equal to the room's sector ID. A becomes B, B becomes C, Z becomes A, and so on. Dashes become spaces.
//
// For example, the real name for qzmt-zixmtkozy-ivhz-343 is very encrypted name.
//
// What is the sector ID of the room where North Pole objects are stored?

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

func main() {
	realRoomMap := parse(loadData("day4.txt"))
	var count int
	for _, v := range realRoomMap {
		count += v
	}
	fmt.Println(count)
	for k, v := range realRoomMap {
		if decrypt(k, v) == "northpole object storage" {
			fmt.Println(v)
			break
		}
	}
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

//Pair pair
type Pair struct {
	Key   string
	Value int
}

//PairList pair list
type PairList []Pair

func (p PairList) Len() int {
	return len(p)
}

func (p PairList) Less(i, j int) bool {
	if p[i].Value != p[j].Value {
		return p[i].Value < p[j].Value
	}

	return p[i].Key > p[j].Key
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

func parse(container []string) map[string]int {
	realRoomMap := make(map[string]int)
	re := regexp.MustCompile("(.*)-([0-9]+)\\[([a-z]{5})\\]")
	for _, line := range container {

		arr := re.FindStringSubmatch(line)
		str := arr[1]

		num, err := strconv.Atoi(arr[2])
		if err != nil {
			log.Fatal(err)
		}

		checkSum := arr[3]
		if isRealRoom(str, checkSum) {
			realRoomMap[str] = num
		}

	}
	return realRoomMap
}

func isRealRoom(roomName string, checkSum string) bool {
	count := make(map[string]int)

	length := len(roomName)
	byt := []byte(roomName)
	for j := 0; j < length; j++ {
		if string(byt[j]) != "-" {
			count[string(byt[j])] = count[string(byt[j])] + 1
		}
	}
	pl := rankByWordCount(count)
	var tmp []string
	for _, k := range pl {
		tmp = append(tmp, k.Key)
	}

	tmpStr := strings.Join(tmp[:5], "")
	return (tmpStr == checkSum) && true
}

func decrypt(realString string, sectorID int) string {
	d := make([]byte, len(realString))
	for i, c := range realString {
		if c == '-' {
			d[i] = ' '
		} else {
			d[i] = byte((int(c-'a')+sectorID)%26 + 'a')
		}
	}
	return string(d)
}
