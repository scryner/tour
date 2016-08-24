// +build OMIT

package main

import "fmt"

type gender int

const (
	male gender = iota
	female
)

type person struct {
	name   string
	age    int
	gender gender
}

func (p person) print() { // value receiver
	fmt.Printf("%s/%v (%d)\n", p.name, p.gender, p.age)
}

func (p *person) aging() { // pointer receiver
	p.age++
}

func main() {
	elsa := person{"Elsa", 21, female}
	anna := person{
		name:   "Anna",
		age:    18,
		gender: female,
	}

	elsa.print()

	anna.aging()
	anna.print()
}
