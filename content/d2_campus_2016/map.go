// +build +OMIT

package main

import "fmt"

func main() {
	arr := []string{"alpha", "bravo", "charlie", "delta", "echo"}

	m := make(map[string]int)
	m["zero"] = 0

	for i, val := range arr {
		m[val] = i + 1
	}

	for key, val := range m {
		fmt.Printf("%d: %s\n", val, key) // 출력 순서 보장 안됨
	}
}
