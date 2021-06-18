package recursion

func Binsearch(arr []int, item int, low int, high int) int {
	if len(arr) <= 1 {
		return 0
	}

	mid := (high + low) / 2
	guess := arr[mid]
	index := mid

	if guess == item {
		return index
	}

	if guess > item {
		high = mid - 1
		return Binsearch(arr, item, low, high)
	} else {
		low = mid + 1
		return Binsearch(arr, item, low, high)
	}
}
