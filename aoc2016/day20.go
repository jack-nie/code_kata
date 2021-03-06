// --- Day 20: Firewall Rules ---
//
// You'd like to set up a small hidden computer here so you can use it to get back into the network later. However, the corporate firewall only allows communication with certain external IP addresses.
//
// You've retrieved the list of blocked IPs from the firewall, but the list seems to be messy and poorly maintained, and it's not clear which IPs are allowed. Also, rather than being written in dot-decimal notation, they are written as plain 32-bit integers, which can have any value from 0 through 4294967295, inclusive.
//
// For example, suppose only the values 0 through 9 were valid, and that you retrieved the following blacklist:
//
// 5-8
// 0-2
// 4-7
// The blacklist specifies ranges of IPs (inclusive of both the start and end value) that are not allowed. Then, the only IPs that this firewall allows are 3 and 9, since those are the only numbers not in any range.
//
// Given the list of blocked IPs you retrieved from the firewall (your puzzle input), what is the lowest-valued IP that is not blocked?
//
// --- Part Two ---
//
// How many IPs are allowed by the blacklist?

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	bans := parse(loadData("day20.txt"))
	fmt.Println(findLowestIp(bans))
	fmt.Println(allowedIps(bans))
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

type bannedIps struct {
	start, end uint32
}

func (b *bannedIps) contains(i uint32) bool {
	return (b.start <= i && b.end >= i)
}

func parse(container []string) []bannedIps {
	var b bannedIps
	var bans []bannedIps
	for _, line := range container {
		startEnd := strings.Split(line, "-")
		start, err := strconv.Atoi(startEnd[0])
		if err != nil {
			log.Fatal(err)
		}
		end, err := strconv.Atoi(startEnd[1])
		if err != nil {
			log.Fatal(err)
		}
		b.start = uint32(start)
		b.end = uint32(end)
		bans = append(bans, b)
	}

	return bans
}

func findLowestIp(bans []bannedIps) uint32 {
HERE:
	for i := uint32(0); i < math.MaxUint32; i++ {
		for _, b := range bans {
			if b.contains(i) {
				continue HERE
			}
		}
		return i
	}
	panic("not found")
}

func allowedIps(bans []bannedIps) int {
	isBanned := func(i uint32) bool {
		for _, b := range bans {
			if b.contains(i) {
				return true
			}
		}
		return false
	}

	n := 0
	for i := uint64(0); i < math.MaxUint32; i++ {
		if isBanned(uint32(i)) {
			continue
		}
		n++
	}
	return n
}
