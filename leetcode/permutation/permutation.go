package main

import "fmt"

func main() {
	result := make([][]int, 0)
	nums := []int{1, 2, 3, 4}
	permutation(nums, []int{}, &result)
	fmt.Println(result)
}

func permutation(nums []int, reserved []int, result *[][]int) {
	if 0 == len(nums) {
		return
	}
	if 1 == len(nums) {
		*result = append(*result, append(reserved, nums[0]))
		return
	}
	permutation(nums[1:], append(reserved, nums[0]), result)
	for i := 1; i < len(nums); i++ {
		nums[i], nums[0] = nums[0], nums[i]
		permutation(nums[1:], append(reserved, nums[0]), result)
		nums[i], nums[0] = nums[0], nums[i]
	}
}
