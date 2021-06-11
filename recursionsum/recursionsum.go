package recursionsum

func Handler(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	return arr[0] + Handler(arr[1:])
}
