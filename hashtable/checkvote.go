package hashtable

import "fmt"

var voted = make(map[string]bool)

func Checkvote(name string) {
	if voted[name] {
		fmt.Printf("Please stop, %s!\n", name)
	} else {
		voted[name] = true
		fmt.Printf("Go ahead, %s!\n", name)
	}
}
