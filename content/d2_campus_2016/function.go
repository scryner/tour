// +build OMIT

package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int64 {
	return int64(a) * int64(b)
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(mul(2, 3))

	/*
    if val, err := div(2, 3); err == nil {
        fmt.Println(val)
    }
	*/
}
