package main

import "fmt"

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 1; j < n-i; j++ {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
}

func bubbleSort2(arr []int) {
	n := len(arr)
	flag := true
	for flag {
		flag = false
		for i := 1; i < n; i++ {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				flag = true
			}
		}
		n--
	}
}

func bubbleSort3(arr []int) {
	flag := len(arr)
	for flag > 0 {
		k := flag
		flag = 0
		for i := 1; i < k; i++ {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				flag = i
			}
		}
	}
}

func main() {
	arr := []int{3, 5, 4, 7, 6, 2, 1}
	bubbleSort(arr)
	fmt.Println(arr)
	arr2 := []int{3, 5, 4, 7, 6, 2, 1}
	bubbleSort2(arr2)
	fmt.Println(arr2)
	arr3 := []int{3, 5, 4, 7, 6, 2, 1}
	bubbleSort3(arr3)
	fmt.Println(arr3)
}
