package main

import "fmt"

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0 && arr[j] > arr[j+1]; j-- {
			arr[j+1], arr[j] = arr[j], arr[j+1]
		}
	}
}

func main() {
	arr := []int{5, 4, 3, 7, 6, 1, 2}
	insertionSort(arr)
	fmt.Println(arr)
}
