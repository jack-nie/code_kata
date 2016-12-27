// --- Day 10: Balance Bots ---
//
// You come upon a factory in which many robots are zooming around handing small microchips to each other.
//
// Upon closer examination, you notice that each bot only proceeds when it has two microchips, and once it does, it gives each one to a different bot or puts it in a marked "output" bin. Sometimes, bots take microchips from "input" bins, too.
//
// Inspecting one of the microchips, it seems like they each contain a single number; the bots must use some logic to decide what to do with each chip. You access the local control computer and download the bots' instructions (your puzzle input).
//
// Some of the instructions specify that a specific-valued microchip should be given to a specific bot; the rest of the instructions indicate what a given bot should do with its lower-value or higher-value chip.
//
// For example, consider the following instructions:
//
// value 5 goes to bot 2
// bot 2 gives low to bot 1 and high to bot 0
// value 3 goes to bot 1
// bot 1 gives low to output 1 and high to bot 0
// bot 0 gives low to output 2 and high to output 0
// value 2 goes to bot 2
// Initially, bot 1 starts with a value-3 chip, and bot 2 starts with a value-2 chip and a value-5 chip.
// Because bot 2 has two microchips, it gives its lower one (2) to bot 1 and its higher one (5) to bot 0.
// Then, bot 1 has two microchips; it puts the value-2 chip in output 1 and gives the value-3 chip to bot 0.
// Finally, bot 0 has two microchips; it puts the 3 in output 2 and the 5 in output 0.
// In the end, output bin 0 contains a value-5 microchip, output bin 1 contains a value-2 microchip, and output bin 2 contains a value-3 microchip. In this configuration, bot number 2 is responsible for comparing value-5 microchips with value-2 microchips.
//
// Based on your instructions, what is the number of the bot that is responsible for comparing value-61 microchips with value-17 microchips?
// --- Part Two ---
//
// What do you get if you multiply together the values of one chip in each of outputs 0, 1, and 2?

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	lines := loadData("day10.txt")
	botMap := make(map[string][]int)
	var done []string
	count := 0
	processInstrcustions(count, lines, botMap, done)
	calculateBotNumber(botMap)
	calculateChipValuesInOutput(botMap)
}

func loadData(filePath string) []string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	var container []string
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		container = append(container, text)
	}
	return container
}

func calculateBotNumber(botMap map[string][]int) {
	re := regexp.MustCompile("[0-9]+")
	for k, v := range botMap {
		if len(v) != 2 {
			continue
		}
		if v[0] == 61 && v[1] == 17 || v[0] == 17 && v[1] == 61 {
			fmt.Println(re.FindString(k))
		}
	}
}

func calculateChipValuesInOutput(botMap map[string][]int) {
	result := 1
	for _, value := range botMap["output0"] {
		result *= value
	}
	for _, value := range botMap["output1"] {
		result *= value
	}
	for _, value := range botMap["output2"] {
		result *= value
	}

	fmt.Println(result)
}

func processInstrcustions(count int, lines []string, botMap map[string][]int, done []string) {
	for count < len(lines) {
		for _, line := range lines {
			if isDone(done, line) {
				continue
			}
			content := strings.Fields(line)
			if len(content) == 6 {
				if _, ok := initialValue(botMap, content); ok {
					count++
					done = append(done, line)
				}
			}
			if len(content) == 12 {
				if _, ok := transferValue(botMap, content); ok {
					count++
					done = append(done, line)
				}
			}
		}
	}
}

func isDone(done []string, line string) bool {
	for _, doneLine := range done {
		if doneLine == line {
			return true
		}
	}
	return false
}

func initialValue(botMap map[string][]int, content []string) (map[string][]int, bool) {
	value, err := strconv.Atoi(content[1])
	if err != nil {
		log.Fatal(err)
	}
	botMap[content[4]+content[5]] = append(botMap[content[4]+content[5]], value)
	return botMap, true
}

func transferValue(botMap map[string][]int, content []string) (map[string][]int, bool) {
	if len(botMap[content[0]+content[1]]) != 2 {
		return botMap, false
	}
	maxValue := max(botMap[content[0]+content[1]])
	minValue := min(botMap[content[0]+content[1]])
	botMap[content[5]+content[6]] = append(botMap[content[5]+content[6]], minValue)
	botMap[content[10]+content[11]] = append(botMap[content[10]+content[11]], maxValue)
	return botMap, true
}

func max(v []int) int {
	if v[0] > v[1] {
		return v[0]
	}
	return v[1]
}

func min(v []int) int {
	if v[0] < v[1] {
		return v[0]
	}
	return v[1]
}
