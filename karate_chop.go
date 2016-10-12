package code_kata

func chop(target int, arr []int) int {
	var left int
	var right int
	var mid int
	left = 0
	right = len(arr) - 1

	mid = left + (right-left)>>1

	for left <= right {
		if arr[mid] < target {
			left = mid + 1
		}

		if arr[mid] > target {
			right = mid - 1
		}

		if arr[mid] == target {
			return mid
		}

		mid = left + (right-left)>>2
	}

	return -1
}
