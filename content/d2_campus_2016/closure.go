// +build OMIT

package main

import (
	"fmt"
	"time"
)

func main() {
	// anonymous function
	i := func() int {
		now := time.Now()
		return int(now.Unix() % 10)
	}()

	fmt.Println(i)

	// closure
	printI := func() {
		fmt.Println("printI:", i)
	}

	printI()

	i = i * 2
	printI()
}
