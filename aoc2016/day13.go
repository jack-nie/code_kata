// --- Day 13: A Maze of Twisty Little Cubicles ---
//
// You arrive at the first floor of this new building to discover a much less welcoming environment than the shiny atrium of the last one. Instead, you are in a maze of twisty little cubicles, all alike.
//
// Every location in this area is addressed by a pair of non-negative integers (x,y). Each such coordinate is either a wall or an open space. You can't move diagonally. The cube maze starts at 0,0 and seems to extend infinitely toward positive x and y; negative values are invalid, as they represent a location outside the building. You are in a small waiting area at 1,1.
//
// While it seems chaotic, a nearby morale-boosting poster explains, the layout is actually quite logical. You can determine whether a given x,y coordinate will be a wall or an open space using a simple system:
//
// Find x*x + 3*x + 2*x*y + y + y*y.
// Add the office designer's favorite number (your puzzle input).
// Find the binary representation of that sum; count the number of bits that are 1.
// If the number of bits that are 1 is even, it's an open space.
// If the number of bits that are 1 is odd, it's a wall.
// For example, if the office designer's favorite number were 10, drawing walls as # and open spaces as ., the corner of the building containing 0,0 would look like this:
//
//   0123456789
// 0 .#.####.##
// 1 ..#..#...#
// 2 #....##...
// 3 ###.#.###.
// 4 .##..#..#.
// 5 ..##....#.
// 6 #...##.###
// Now, suppose you wanted to reach 7,4. The shortest route you could take is marked as O:
//
//   0123456789
// 0 .#.####.##
// 1 .O#..#...#
// 2 #OOO.##...
// 3 ###O#.###.
// 4 .##OO#OO#.
// 5 ..##OOO.#.
// 6 #...##.###
// Thus, reaching 7,4 would take a minimum of 11 steps (starting from your current location, 1,1).
//
// What is the fewest number of steps required for you to reach 31,39?
// --- Part Two ---
//
// How many locations (distinct x,y coordinates, including your starting location) can you reach in at most 50 steps?

package main

import (
	"fmt"
	"strconv"
)

const (
	input = 1358
)

func main() {
	start := make(map[pos]int)
	mazeOuter := make(map[pos]int)
	start[pos{1, 1}] = 1
	count1 := 0
	count2 := 0

	for !contains(start) {
		count1++
		start = step(start, mazeOuter)
		if count1 == 50 {
			count2 = len(mazeOuter)
		}
	}
	fmt.Println(count1)
	fmt.Println(count2)
}

func contains(start map[pos]int) bool {
	var flag bool
	for k := range start {
		if !(k.x == 31 && k.y == 39) {
			continue
		} else {
			flag = true
		}
	}
	return flag
}

type pos struct {
	x, y int
}

func isEven(x, y int) bool {
	num := x*x + 3*x + 2*x*y + y + y*y
	binInt := strconv.FormatInt(int64(num+input), 2)
	return isEvenInt(binInt)
}

func isEvenInt(s string) bool {
	var i int
	for _, c := range s {
		if c == '1' {
			i++
		}
	}
	return (i%2 == 0)
}

func step(start map[pos]int, mazeOuter map[pos]int) map[pos]int {
	mazeInner := make(map[pos]int)
	for k := range start {
		x := k.x
		y := k.y
		if mazeInner[pos{x, y + 1}] != 1 && isEven(x, y+1) {
			mazeInner[pos{x, y + 1}] = 1
			mazeOuter[pos{x, y + 1}] = 1
		}
		if mazeInner[pos{x + 1, y}] != 1 && isEven(x+1, y) {
			mazeInner[pos{x + 1, y}] = 1
			mazeOuter[pos{x + 1, y}] = 1
		}
		if y != 0 && mazeInner[pos{x, y - 1}] != 1 && isEven(x, y-1) {
			mazeInner[pos{x, y - 1}] = 1
			mazeOuter[pos{x, y - 1}] = 1
		}
		if x != 0 && mazeInner[pos{x - 1, y}] != 1 && isEven(x-1, y) {
			mazeInner[pos{x - 1, y}] = 1
			mazeOuter[pos{x - 1, y}] = 1
		}
	}
	return mazeInner
}
