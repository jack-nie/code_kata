// --- Day 8: Two-Factor Authentication ---
//
// You come across a door implementing what you can only assume is an implementation of two-factor authentication after a long game of requirements telephone.
//
// To get past the door, you first swipe a keycard (no problem; there was one on a nearby desk). Then, it displays a code on a little screen, and you type that code on a keypad. Then, presumably, the door unlocks.
//
// Unfortunately, the screen has been smashed. After a few minutes, you've taken everything apart and figured out how it works. Now you just have to work out what the screen would have displayed.
//
// The magnetic strip on the card you swiped encodes a series of instructions for the screen; these instructions are your puzzle input. The screen is 50 pixels wide and 6 pixels tall, all of which start off, and is capable of three somewhat peculiar operations:
//
// rect AxB turns on all of the pixels in a rectangle at the top-left of the screen which is A wide and B tall.
// rotate row y=A by B shifts all of the pixels in row A (0 is the top row) right by B pixels. Pixels that would fall off the right end appear at the left end of the row.
// rotate column x=A by B shifts all of the pixels in column A (0 is the left column) down by B pixels. Pixels that would fall off the bottom appear at the top of the column.
// For example, here is a simple sequence on a smaller screen:
//
// rect 3x2 creates a small rectangle in the top-left corner:
//
// ###....
// ###....
// .......
// rotate column x=1 by 1 rotates the second column down by one pixel:
//
// #.#....
// ###....
// .#.....
// rotate row y=0 by 4 rotates the top row right by four pixels:
//
// ....#.#
// ###....
// .#.....
// rotate column x=1 by 1 again rotates the second column down by one pixel, causing the bottom pixel to wrap back to the top:
//
// .#..#.#
// #.#....
// .#.....
// As you can see, this display technology is extremely powerful, and will soon dominate the tiny-code-displaying-screen market. That's what the advertisement on the back of the display tries to convince you, anyway.
//
// There seems to be an intermediate check of the voltage used by the display: after you swipe your card, if the screen did work, how many pixels should be lit?

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	keypad := Keypad(50, 6)
	strs := loadData("day8.txt")
	for _, str := range strs {
		instruction := strings.Fields(str)
		switch instruction[0] {
		case "rect":
			fillInPads(instruction[1], keypad)
		case "rotate":
			move(instruction[1:], keypad)
		}
	}

	count := calculatePad(keypad)
	fmt.Println(count)
	printPad(keypad)
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

func fillInPads(instructions string, keypad [][]bool) {
	area := strings.Split(instructions, "x")
	x, err := strconv.Atoi(area[0])
	if err != nil {
		log.Fatal(err)
	}
	y, err := strconv.Atoi(area[1])
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			keypad[i][j] = true
		}
	}
}

func move(instructions []string, keypad [][]bool) {
	direction := strings.Split(instructions[1], "=")

	position, err := strconv.Atoi(direction[1])
	if err != nil {
		log.Fatal(err)
	}
	step, err := strconv.Atoi(instructions[3])
	if err != nil {
		log.Fatal(err)
	}

	switch direction[0] {
	case "x":
		moveColumn(position, step, keypad)
	case "y":
		moveRow(position, step, keypad)
	}
}

func moveColumn(position int, step int, keypad [][]bool) {
	var tmp []bool
	for i := 0; i < 6; i++ {
		tmp = append(tmp, keypad[i][position])
	}

	for j := 0; j < 6; j++ {
		realPosition := (j + step) % 6
		keypad[realPosition][position] = tmp[j]
	}
}

func moveRow(position int, step int, keypad [][]bool) {
	var tmp []bool
	for i := 0; i < 50; i++ {
		tmp = append(tmp, keypad[position][i])
	}

	for j := 0; j < 50; j++ {
		realPosition := (j + step) % 50
		keypad[position][realPosition] = tmp[j]
	}
}

func calculatePad(keypad [][]bool) int {
	count := 0
	for i := 0; i < 6; i++ {
		for j := 0; j < 50; j++ {
			if keypad[i][j] {
				count++
			}
		}
	}

	return count
}

func Keypad(x, y int) [][]bool {
	keypad := make([][]bool, y)
	for i := range keypad {
		keypad[i] = make([]bool, x)
		for j := range keypad[i] {
			keypad[i][j] = false
		}
	}
	return keypad
}

func printPad(keypad [][]bool) {
	for i := 0; i < 6; i++ {
		for j := 0; j < 50; j++ {
			if keypad[i][j] {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}
