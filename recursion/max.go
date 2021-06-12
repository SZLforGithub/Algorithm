package recursion

func Max(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	if len(arr) == 1 {
		return arr[0]
	}

	max := Max(arr[1:])

	if arr[0] > max {
		max = arr[0]
	}

	return max
}
