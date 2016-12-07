// --- Day 7: Internet Protocol Version 7 ---
//
// While snooping around the local network of EBHQ, you compile a list of IP addresses (they're IPv7, of course; IPv6 is much too limited). You'd like to figure out which IPs support TLS (transport-layer snooping).
//
// An IP supports TLS if it has an Autonomous Bridge Bypass Annotation, or ABBA. An ABBA is any four-character sequence which consists of a pair of two different characters followed by the reverse of that pair, such as xyyx or abba. However, the IP also must not have an ABBA within any hypernet sequences, which are contained by square brackets.
//
// For example:
//
// abba[mnop]qrst supports TLS (abba outside square brackets).
// abcd[bddb]xyyx does not support TLS (bddb is within square brackets, even though xyyx is outside square brackets).
// aaaa[qwer]tyui does not support TLS (aaaa is invalid; the interior characters must be different).
// ioxxoj[asdfgh]zxcvbn supports TLS (oxxo is outside square brackets, even though it's within a larger string).
// How many IPs in your puzzle input support TLS?

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
)

func main() {
	counter := 0
	dat := loadData("day7.txt")
	for _, strs := range dat {
		if allABBA(strs) {
			counter++
		}
	}
	fmt.Println(counter)
}

func allABBA(strs string) bool {
	var flag bool
	f := func(c rune) bool {
		return string(c) == "[" || string(c) == "]"
	}
	s := strings.FieldsFunc(strs, f)

	for i, str := range s {
		if i%2 == 0 && isABBA(str) {
			flag = true
		}

		if i%2 != 0 && isABBA(str) {
			fmt.Println(str)
			return false
		}

	}
	return flag
}

func isABBA(str string) bool {
	byt := []byte(str)
	for i := 0; i < len(byt)-3; i++ {
		if reflect.DeepEqual(byt[i:i+4], reverseSlice(byt[i:i+4])) && !isSame(byt[i:i+4]) {
			return true
		}
	}
	return false
}

func reverseSlice(str []byte) []byte {
	s := make([]byte, len(str))
	copy(s, str)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func isSame(str []byte) bool {
	return str[0] == str[1] &&
		str[1] == str[2] &&
		str[2] == str[3] &&
		true
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
