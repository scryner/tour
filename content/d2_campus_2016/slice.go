// +build OMIT

package main

import "fmt"

func main() {
	stringSet := []string{"hello", "world",
		"d2", "campus seminar"}

	foreachPrint := func(arr []string) {
		// for-each loop
		for _, entry := range arr {
			fmt.Println(entry)
		}
	}

	fmt.Println("set length:", len(stringSet))
	foreachPrint(stringSet)

	// append an element
	stringSet = append(stringSet, "naver")
	foreachPrint(stringSet)

	// slicing slices
	subSet1 := stringSet[1:3]
	foreachPrint(subSet1)

	subSet2 := stringSet[2:]
	foreachPrint(subSet2)

	subSet3 := stringSet[:3]
	foreachPrint(subSet3)
}
