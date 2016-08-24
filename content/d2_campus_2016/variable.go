// +build OMIT

package main

import "fmt"

func main() {
	// 변수 선언
	var a int = 10
	var b string = "hello, world"
	var c bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// 짧은 변수 선언
	d := 10.0
	e := true
	f := "hello, world"

	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	// pointer 형
	var i *int = nil
	i = &a

	fmt.Println(*i)
	a = 20
	fmt.Println(*i)
}
