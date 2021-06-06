package main

import (
	"fmt"

	// binSearch "github.com/SZLforGithub/Algorithm/binarysearch"
	selectSearch "github.com/SZLforGithub/Algorithm/selectionsearch"
)

func main() {
	// myList := []int{1, 3, 5, 7, 9, 11}
	// fmt.Println(binSearch.Handler(myList, 3))  // return 1
	// fmt.Println(binSearch.Handler(myList, -1)) // return -1

	myList := []int{4, 1, 7, 0, 2, 8, 3, 10, 1}
	fmt.Println(selectSearch.Handler(myList))
}
