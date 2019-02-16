package main

import "fmt"

func shellSort(arr []int) {
	if len(arr) == 0 {
		return
	}
	n := len(arr)
	for gap := n / 2; gap > 0; gap /= 2 {
		for i := gap; i < n; i++ {
			for j := i - gap; j >= 0 && arr[j] > arr[j+gap]; j -= gap {
				arr[j], arr[j+gap] = arr[j+gap], arr[j]
			}
		}
	}
}

func main() {
	arr := []int{3, 6, 5, 4, 8, 7, 2, 1}
	shellSort(arr)
	fmt.Println(arr)
}
