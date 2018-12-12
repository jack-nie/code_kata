package main

import "fmt"

func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{[]int{}}
	}
	result := [][]int{}
	backtracking(nums, &result, []int{}, 0)
	return result
}

func backtracking(nums []int, result *[][]int, tempList []int, start int) {
	*result = append(*result, tempList)
	for i := start; i < len(nums); i++ {
		tempList = append(tempList, nums[i])
		backtracking(nums, result, tempList, i+1)
		tempList = tempList[: len(tempList)-1 : len(tempList)-1]
	}
}

func main() {
	result := subsets([]int{1, 2, 3})
	fmt.Println(result)
}
