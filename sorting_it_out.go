package code_kata

func quickSort(arr []int, low int, high int) {
	var pivot, j, tmp, i int

	if low < high {
		pivot = low
		i = low
		j = high

		for i < j {
			for arr[i] <= arr[pivot] && i < high {
				i++
			}

			for arr[j] > arr[pivot] {
				j--
			}

			if i < j {
				tmp = arr[i]
				arr[i] = arr[j]
				arr[j] = tmp
			}
		}

		tmp = arr[pivot]
		arr[pivot] = arr[j]
		arr[j] = tmp
		quickSort(arr, low, j-1)
		quickSort(arr, j+1, high)
	}
}
