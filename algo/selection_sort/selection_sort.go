package main

import "fmt"

func selectionSort(arr []int) {
	if len(arr) == 0 {
		return
	}

	for i := 0; i < len(arr); i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				arr[minIndex], arr[j] = arr[j], arr[minIndex]
			}
		}
	}
}

func main() {
	arr := []int{5, 4, 3, 7, 6, 2, 1}
	selectionSort(arr)
	fmt.Println(arr)
}
