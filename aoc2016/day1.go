// --- Day 1: No Time for a Taxicab ---
//
// Santa's sleigh uses a very high-precision clock to guide its movements, and the clock's oscillator is regulated by stars. Unfortunately, the stars have been stolen... by the Easter Bunny. To save Christmas, Santa needs you to retrieve all fifty stars by December 25th.
//
// Collect stars by solving puzzles. Two puzzles will be made available on each day in the advent calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!
//
// You're airdropped near Easter Bunny Headquarters in a city somewhere. "Near", unfortunately, is as close as you can get - the instructions on the Easter Bunny Recruiting Document the Elves intercepted start here, and nobody had time to work them out further.
//
// The Document indicates that you should start at the given coordinates (where you just landed) and face North. Then, follow the provided sequence: either turn left (L) or right (R) 90 degrees, then walk forward the given number of blocks, ending at a new intersection.
//
// There's no time to follow such ridiculous instructions on foot, though, so you take a moment and work out the destination. Given that you can only walk on the street grid of the city, how far is the shortest path to the destination?
//
// For example:
//
// Following R2, L3 leaves you 2 blocks East and 3 blocks North, or 5 blocks away.
// R2, R2, R2 leaves you 2 blocks due South of your starting position, which is 2 blocks away.
// R5, L5, R5, R3 leaves you 12 blocks away.
// How many blocks away is Easter Bunny HQ?
//
// --- Part Two ---
//
// Then, you notice the instructions continue on the back of the Recruiting Document. Easter Bunny HQ is actually at the first location you visit twice.
//
// For example, if your instructions are R8, R4, R4, R8, the first location you visit twice is 4 blocks away, due East.
//
// How many blocks away is the first location you visit twice?

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type pos struct {
	x, y int
}

type dir struct {
	left   bool
	length int
}

const (
	north = iota
	west
	south
	east
)

func walk(p *pos, orientation int, length int) {
	switch orientation {
	case north:
		p.y += length
	case south:
		p.y -= length
	case east:
		p.x += length
	case west:
		p.x -= length
	}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (p *pos) dist() int {
	return Abs(p.x) + Abs(p.y)
}

func parse(filePath string) []dir {
	var dirs []dir

	s := loadData(filePath)
	c := strings.Split(s[0], ", ")
	for _, d := range c {

		step, err := strconv.Atoi(d[1:])
		if err != nil {
			log.Fatal(err)
		}

		dirs = append(dirs, dir{
			left:   d[0] == 'L',
			length: step,
		})
	}
	return dirs
}

func follow(dirs []dir) pos {
	var p pos
	o := north
	for _, d := range dirs {
		if d.left {
			o = (o + 1) % 4
		} else {
			o = (o + 3) % 4
		}
		walk(&p, o, d.length)
	}
	return p
}

func visitTwice(dirs []dir) pos {
	visited := make(map[pos]bool)
	var p pos
	o := north
	visited[p] = true
	for _, d := range dirs {
		if d.left {
			o = (o + 1) % 4
		} else {
			o = (o + 3) % 4
		}

		for i := 0; i < d.length; i++ {
			walk(&p, o, 1)
			if visited[p] {
				return p
			}
			visited[p] = true
		}
	}
	return p
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

func main() {
	dirs := parse("day1.txt")
	end := follow(dirs)
	fmt.Println(end.dist())

	end = visitTwice(dirs)
	fmt.Println(end.dist())
}
