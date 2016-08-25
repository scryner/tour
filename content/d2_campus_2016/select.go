// +build OMIT

package main

import (
	"fmt"
	"time"
)

const answerToEverything = 42

func getAnswerToEverything() chan int {
	ch := make(chan int)

	go func() {
		// 0 ~ 9초간 슬립
		nowUnix := time.Now().Unix()
		time.Sleep(time.Second *
			time.Duration(nowUnix%10))
		ch <- answerToEverything
	}()

	return ch
}

// time.After() 함수와 같은 일을 수행
func after(duration time.Duration) chan bool {
	ch := make(chan bool)

	go func() {
		time.Sleep(duration)
		ch <- true
	}()

	return ch
}

func main() {
	fmt.Println(
		"Calculating the answer to everything")

	select {
	case result := <-getAnswerToEverything():
		fmt.Println(
			"I just found the answer to everything:",
			result)
	case <-after(time.Second * 3):
		fmt.Println("ERROR: timeouted!")
	}
}
