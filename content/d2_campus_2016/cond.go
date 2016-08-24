// +build OMIT

package main

import "fmt"

func main() {
	a := 100
	b := "hello"

	// if 문
	if a > 100 {
		fmt.Println("greater than 100")
	}

	if b == "hello" {
		fmt.Println(b, "world")
	} else {
        fmt.Println("hello", b)
    }

	// for문; C for-loop 스타일
	for i := 0; i < 3; i++ {
		fmt.Printf("i: %d\n", i)
	}

	// for문; C while-loop 스타일
	j := 0
	for j < 3 {
		fmt.Printf("j: %d\n", j)
		j++
	}
	
	/* for문; 무한 루프
	for {
		fmt.Println("in infinite loop")
	}
	*/

	// switch 문; C switch-case 스타일
	switch b {
	case "hello":
		fmt.Println(b, "world")
	case "world":
		fmt.Println("hello", b)
	case "안녕", "世界":
		fmt.Println("안녕, 世界")
	default:
		fmt.Println("hello world")
	}

	// switch 문; if-then-else 스타일
	switch {
	case a < 100:
		fmt.Println("less than 100")
	case a > 100:
		fmt.Println("greater than 100")
	case a == 100:
		fmt.Println("equal to 100")
	case b == "hello":
		fmt.Println("true, but never reached")
	default:
		fmt.Println("never reached")
	}
}
