package selectionsearch

func Handler(arr []int) []int {
	var newArr []int
	length := len(arr)

	for i := 0; i < length; i++ {
		smallestIndex := findSmallest(arr)
		newArr = append(newArr, arr[smallestIndex])
		arr = remove(arr, smallestIndex)
	}

	return newArr
}

func findSmallest(arr []int) int {
	smallestIndex := 0
	smallest := arr[smallestIndex]

	for i, v := range arr {
		if v < smallest {
			smallestIndex = i
			smallest = v
		}
	}

	return smallestIndex
}

func remove(arr []int, index int) []int {
	return append(arr[:index], arr[index+1:]...)
}
