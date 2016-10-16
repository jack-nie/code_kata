package code_kata

func division(n int) int {
	if n == 1 {
		return 2
	}

	if n == 2 {
		return 3
	}

	if n > 2 {
		return division(n-1) + division(n-2)
	}

	return -1
}
