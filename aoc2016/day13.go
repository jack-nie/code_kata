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

package main

import (
	"fmt"
	"log"
	"strconv"
)

const (
	endX   = 31
	endY   = 39
	input  = 1358
	length = 100
)

func main() {
	maze := make(map[pos]int)
	count := 0

	step(0, 0, maze)
	for _, _ = range maze {
		count++
	}

	fmt.Println(count)
	fmt.Println(maze)
}

type pos struct {
	x, y int
}

func isOdd(x, y int) bool {
	num := x*x + 3*x + 2*x*y + y + y*y
	binInt, err := strconv.Atoi(strconv.FormatInt(int64(num+input), 2))
	if err != nil {
		log.Fatal(err)
	}
	return isOddInt(binInt)
}

func isOddInt(n int) bool {
	n -= (n >> 1) & 0x5555555555555555
	n = (n>>2)&0x3333333333333333 + n&0x3333333333333333
	n += n >> 4
	n &= 0x0f0f0f0f0f0f0f0f
	n *= 0x0101010101010101
	return byte(n>>56)%2 != 0
}

func step(x, y int, maze map[pos]int) {
	maze[pos{x, y}] = 1
	if x == endX && y == endY {
		return
	}

	fmt.Println(maze)
	if y < (length-1) && isOdd(x, y+1) {
		step(x, y+1, maze)
	}
	if x < (length-1) && isOdd(x+1, y) {
		step(x+1, y, maze)
	}
	if y > 0 && isOdd(x, y-1) {
		step(x, y-1, maze)
	}
	if x > 0 && isOdd(x-1, y) {
		step(x-1, y, maze)
	}
}
