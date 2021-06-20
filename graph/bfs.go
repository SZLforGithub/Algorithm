package graph

import (
	"fmt"
)

func Bfs(graph map[string][]string, name string) bool {
	var searchQueue []string
	searchQueue = append(searchQueue, graph[name]...)
	var searched []string

	for len(searchQueue) > 0 {
		var person = searchQueue[0]
		searchQueue = searchQueue[1:]

		if isStringIn(person, searched) {
			continue
		}

		if isSeller(person) {
			fmt.Println(person, "is a mango seller!")
			return true
		}

		searched = append(searched, person)
		searchQueue = append(searchQueue, graph[person]...)
	}

	return false
}

func isSeller(name string) bool {
	return string(name[len(name)-1]) == "m"
}

func isStringIn(name string, list []string) bool {
	for _, v := range list {
		if v == name {
			return true
		}
	}

	return false
}
