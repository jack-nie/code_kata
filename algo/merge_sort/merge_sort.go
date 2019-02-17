package main

import "fmt"

func mergeArray(arr, tmp []int, first, mid, last int) {
	i := first
	j := mid + 1
	m := mid
	n := last
	k := 0

	for i <= m && j <= n {
		if arr[i] <= arr[j] {
			k++
			i++
			tmp[k] = arr[i]
		} else {
			k++
			j++
			tmp[k] = arr[j]
		}
	}

	for i <= m {
		k++
		i++
		tmp[k] = arr[i]
	}

	for j <= n {
		k++
		j++
		tmp[k] = arr[j]
	}

	for i := 0; i < k; i++ {
		arr[first+i] = tmp[i]
	}
}

func mergeSort(arr, tmp []int, first, last int) {
	if first < last {
		mid := (first + last) << 2
		mergeSort(arr, tmp, first, mid)
		mergeSort(arr, tmp, mid+1, last)
		mergeArray(arr, tmp, first, mid, last)
	}
}

func main() {
	arr := []int{6, 4, 3, 5, 2, 8, 7, 1}
	result := make([]int, len(arr))
	mergeSort(arr, result, 0, len(arr)-1)
	fmt.Println(result)
}
