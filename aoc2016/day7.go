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
// --- Part Two ---
//
// You would also like to know which IPs support SSL (super-secret listening).
//
// An IP supports SSL if it has an Area-Broadcast Accessor, or ABA, anywhere in the supernet sequences (outside any square bracketed sections), and a corresponding Byte Allocation Block, or BAB, anywhere in the hypernet sequences. An ABA is any three-character sequence which consists of the same character twice with a different character between them, such as xyx or aba. A corresponding BAB is the same characters but in reversed positions: yxy and bab, respectively.
//
// For example:
//
// aba[bab]xyz supports SSL (aba outside square brackets with corresponding bab within square brackets).
// xyx[xyx]xyx does not support SSL (xyx, but no corresponding yxy).
// aaa[kek]eke supports SSL (eke in supernet with corresponding kek in hypernet; the aaa sequence is not related, because the interior character must be different).
// zazbz[bzb]cdb supports SSL (zaz has no corresponding aza, but zbz has a corresponding bzb, even though zaz and zbz overlap).
// How many IPs in your puzzle input support SSL?

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
	counter2 := 0
	dat := loadData("day7.txt")
	for _, strs := range dat {
		if allABBA(strs) {
			counter++
		}
	}

	for _, strs := range dat {
		if isABA(strs) {
			counter2++
		}
	}
	fmt.Println(counter)
	fmt.Println(counter2)
}

func allABBA(strs string) bool {
	var flag bool
	s := parseStr(strs)

	for i, str := range s {
		if i%2 == 0 && isABBA(str) {
			flag = true
		}

		if i%2 != 0 && isABBA(str) {
			return false
		}

	}
	return flag
}

func parseStr(strs string) []string {
	f := func(c rune) bool {
		return string(c) == "[" || string(c) == "]"
	}
	return strings.FieldsFunc(strs, f)
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
	defer dat.Close()
	scanner := bufio.NewScanner(dat)
	for scanner.Scan() {
		text := scanner.Text()
		container = append(container, text)
	}
	return container
}

func collectABA(str string, abas [][]byte) [][]byte {
	byt := []byte(str)
	for i := 0; i < len(byt)-2; i++ {
		if byt[i] == byt[i+2] && byt[i] != byt[i+1] {
			abas = append(abas, byt[i:i+3])
		}
	}
	return abas
}

func contains(source [][]byte, target []byte) bool {
	for _, item := range source {
		if reflect.DeepEqual(item, target) {
			return true
		}
	}
	return false
}

func isABA(strs string) bool {
	s := parseStr(strs)
	var abas [][]byte
	var babs [][]byte

	for i, item := range s {
		if i%2 == 0 {
			abas = collectABA(item, abas)
		}

		if i%2 != 0 {
			babs = collectABA(item, babs)
		}
	}

	for _, item := range abas {
		var tmp = []byte{item[1], item[0], item[1]}
		if contains(babs, tmp) {
			return true
		}
	}

	return false
}
