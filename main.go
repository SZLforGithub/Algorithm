package main

import (
	"fmt"

	binSearch "github.com/SZLforGithub/Algorithm/binarysearch"
	sum "github.com/SZLforGithub/Algorithm/recursionsum"
	selectSearch "github.com/SZLforGithub/Algorithm/selectionsearch"
)

func main() {
	showRecursionSum()
}

func showBinSearch() {
	myList := []int{1, 3, 5, 7, 9, 11}
	fmt.Println(binSearch.Handler(myList, 3))  // return 1
	fmt.Println(binSearch.Handler(myList, 12)) // return -1 if not found
}

func showSelectSearch() {
	myList := []int{4, 1, 7, 0, 2, 8, 3, 10, 1}
	fmt.Println(selectSearch.Handler(myList))
}

func showRecursionSum() {
	myList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(sum.Handler(myList))
}
