package main

import "fmt"

func binarySearch(arr []int, target, low, high int) bool {
	if len(arr) == 0 {
		return false
	}
	for low < high {
		middle := (low + high) >> 1
		if arr[middle] == target {
			return true
		} else if arr[middle] < target {
			low = middle + 1
		} else {
			high = middle - 1
		}
	}
	return false
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(binarySearch(arr, 12, 0, 5))
}
