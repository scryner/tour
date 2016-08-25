// +build OMIT

package main

import "fmt"

func main() {
	c := make(chan int)
	done := make(chan int)

	go func() {
		// producer
		for i := 0; i < 5; i++ {
			c <- i
		}

		close(c)
	}()

	go func() {
		// consumer
		for i := range c {
			fmt.Println(i)
		}

		done <- 1
	}()

	<-done
	fmt.Println("all done")
}
