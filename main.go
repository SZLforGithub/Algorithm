package main

import (
	"fmt"

	binSearch "github.com/SZLforGithub/Algorithm/binarysearch"
	hashTable "github.com/SZLforGithub/Algorithm/hashtable"
	quickSort "github.com/SZLforGithub/Algorithm/quicksort"
	recursion "github.com/SZLforGithub/Algorithm/recursion"
	selectSearch "github.com/SZLforGithub/Algorithm/selectionsearch"
)

func main() {
	showVoteCheck()
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
	fmt.Println(recursion.Sum(myList))
}

func showRecursionCount() {
	myList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(recursion.Count(myList))
}

func showRecursionMax() {
	myList := []int{1, 2, 3, 4, 5, 11, 6, 7, 8, 9, 10}
	fmt.Println(recursion.Max(myList))
}

func showRecursionBinarySearch() {
	myList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	fmt.Println(recursion.Binsearch(myList, 9, 0, 10))
}

func showQuickSort() {
	myList := []int{4, 1, 7, 0, 2, 8, 3, 10, 1}
	fmt.Println(quickSort.Handler(myList))
}

func showHashtable() {
	fmt.Println(hashTable.Example())
}

func showVoteCheck() {
	hashTable.Checkvote("Louis")
	hashTable.Checkvote("Jesslyn")
	hashTable.Checkvote("Louis")
}
