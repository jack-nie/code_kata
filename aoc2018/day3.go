// --- Day 3: No Matter How You Slice It ---
// The Elves managed to locate the chimney-squeeze prototype fabric for Santa's suit (thanks to someone who helpfully wrote its box IDs on the wall of the warehouse in the middle of the night). Unfortunately, anomalies are still affecting them - nobody can even agree on how to cut the fabric.
//
// The whole piece of fabric they're working on is a very large square - at least 1000 inches on each side.
//
// Each Elf has made a claim about which area of fabric would be ideal for Santa's suit. All claims have an ID and consist of a single rectangle with edges parallel to the edges of the fabric. Each claim's rectangle is defined as follows:
//
// The number of inches between the left edge of the fabric and the left edge of the rectangle.
// The number of inches between the top edge of the fabric and the top edge of the rectangle.
// The width of the rectangle in inches.
// The height of the rectangle in inches.
// A claim like #123 @ 3,2: 5x4 means that claim ID 123 specifies a rectangle 3 inches from the left edge, 2 inches from the top edge, 5 inches wide, and 4 inches tall. Visually, it claims the square inches of fabric represented by # (and ignores the square inches of fabric represented by .) in the diagram below:
//
// ...........
// ...........
// ...#####...
// ...#####...
// ...#####...
// ...#####...
// ...........
// ...........
// ...........
// The problem is that many of the claims overlap, causing two or more claims to cover part of the same areas. For example, consider the following claims:
//
// #1 @ 1,3: 4x4
// #2 @ 3,1: 4x4
// #3 @ 5,5: 2x2
// Visually, these claim the following areas:
//
// ........
// ...2222.
// ...2222.
// .11XX22.
// .11XX22.
// .111133.
// .111133.
// ........
// The four square inches marked with X are claimed by both 1 and 2. (Claim 3, while adjacent to the others, does not overlap either of them.)
//
// If the Elves all proceed with their own plans, none of them will have enough fabric. How many square inches of fabric are within two or more claims?

package aoc

import (
	"regexp"
	"strconv"
)

type coord struct {
	l int
	t int
}

func mapClaims(data []string) (map[coord][]int, map[int][]int) {
	m := make(map[coord][]int)
	claims := make([][]int, len(data))
	overlaps := make(map[int][]int)
	r, _ := regexp.Compile("-?[0-9]+")
	for i, d := range data {
		claimsVal := r.FindAllString(d, -1)
		for _, valStr := range claimsVal {
			val, _ := strconv.Atoi(valStr)
			claims[i] = append(claims[i], val)
		}
		id := claims[i][0]
		startX := claims[i][1]
		startY := claims[i][2]
		width := claims[i][3]
		height := claims[i][4]
		overlaps[id] = []int{}
		for l := startX; l < startX+width; l++ {
			for t := startY; t < startY+height; t++ {
				claimSet := m[coord{l, t}]
				for _, number := range claimSet {
					overlaps[number] = append(overlaps[number], id)
					overlaps[id] = append(overlaps[id], number)
				}
				claimSet = append(claimSet, id)
				m[coord{l, t}] = claimSet
			}
		}

	}
	return m, overlaps
}

func getLenOverlaps() int {
	data := loadData("day3_2.txt")

	m, _ := mapClaims(data)

	partA := 0
	for _, v := range m {
		if len(v) >= 2 {
			partA++
		}
	}
	return partA
}

func getLenOverlapsTwo() int {
	data := loadData("day3_2.txt")
	_, overlaps := mapClaims(data)
	partB := []int{}
	for k, v := range overlaps {
		if len(v) == 0 {
			partB = append(partB, k)
		}
	}
	return partB[0]
}
