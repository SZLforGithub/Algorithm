package quicksort

func Handler(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	var less []int
	var greater []int

	for _, v := range arr[1:] {
		if v < pivot {
			less = append(less, v)
		} else {
			greater = append(greater, v)
		}
	}

	less = append(Handler(less), pivot)
	greater = Handler(greater)
	return append(less, greater...)
}
