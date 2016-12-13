// --- Day 3: Squares With Three Sides ---
//
// Now that you can think clearly, you move deeper into the labyrinth of hallways and office furniture that makes up this part of Easter Bunny HQ. This must be a graphic design department; the walls are covered in specifications for triangles.
//
// Or are they?
//
// The design document gives the side lengths of each triangle it describes, but... 5 10 25? Some of these aren't triangles. You can't help but mark the impossible ones.
//
// In a valid triangle, the sum of any two sides must be larger than the remaining side. For example, the "triangle" given above is impossible, because 5 + 10 is not larger than 25.
//
// In your puzzle input, how many of the listed triangles are possible?

// --- Part Two ---
//
// Now that you've helpfully marked up their design documents, it occurs to you that triangles are specified in groups of three vertically. Each set of three numbers in a column specifies a triangle. Rows are unrelated.
//
// For example, given the following specification, numbers with the same hundreds digit would be part of the same triangle:
//
// 101 301 501
// 102 302 502
// 103 303 503
// 201 401 601
// 202 402 602
// 203 403 603
// In your puzzle input, and instead reading by columns, how many of the listed triangles are possible?

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
	count := 0
	count2 := 0
	dat := loadData("day3.txt")
	for _, nums := range dat {
		res0, res1, res2 := parseNums(nums)
		if validTriangles(res0, res1, res2) {
			count++
		}
	}

	dat2 := parseNumsTwo(dat)

	for _, nums := range dat2 {
		res0, res1, res2 := parseNums(nums)
		if validTriangles(res0, res1, res2) {
			count2++
		}
	}

	fmt.Println(count)
	fmt.Println(count2)
}

func loadData(filePath string) [][]string {
	dat, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	var container [][]string
	scanner := bufio.NewScanner(dat)

	for scanner.Scan() {
		text := strings.Fields(scanner.Text())
		container = append(container, text)
	}

	return container
}

func validTriangles(res0 int, res1 int, res2 int) bool {
	return true && res0 > 0 && res1 > 0 && res2 > 0 && (res0+res1 > res2) && (res0+res2 > res1) && (res1+res2 > res0)
}

func parseNums(nums []string) (int, int, int) {
	res0, err := strconv.Atoi(nums[0])
	if err != nil {
		log.Fatal(err)
	}
	res1, err := strconv.Atoi(nums[1])
	if err != nil {
		log.Fatal(err)
	}
	res2, err := strconv.Atoi(nums[2])
	if err != nil {
		log.Fatal(err)
	}
	return res0, res1, res2
}

func parseNumsTwo(dat [][]string) [][]string {
	length := len(dat)
	var container [][]string
	for i := 0; i < length; i += 3 {
		res0 := dat[i+0]
		res1 := dat[i+1]
		res2 := dat[i+2]
		for j := 0; j < 3; j++ {
			var innerContainer []string
			innerContainer = append(innerContainer, res0[j], res1[j], res2[j])
			container = append(container, innerContainer)
		}
	}
	return container
}
