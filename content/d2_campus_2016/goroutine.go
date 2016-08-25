// +build OMIT

package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 3; i++ {
		go func(id int) {
			fmt.Println("goroutine", id)
		}(i)
	}

	time.Sleep(time.Second * 1)
}
