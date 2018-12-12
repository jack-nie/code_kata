package main

import "fmt"

func main() {
	arr := []int{2, 4, 3, 7, 6, 9}
	quickSort(arr, 0, 4)
	fmt.Println(arr)
}

func partition(arr []int, left, right int) int {
	pivot := arr[right]
	i := left - 1
	for j := left; j < right; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[right] = arr[right], arr[i+1]
	return i + 1
}

func quickSort(arr []int, left, right int) {
	if len(arr) == 0 {
		return
	}
	if left < right {
		pivot := partition(arr, left, right)
		quickSort(arr, left, pivot-1)
		quickSort(arr, pivot+1, right)
	}
}
