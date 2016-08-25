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

func main() {
	elsa := person{"Elsa", 21, female}
	elsa.print()
}

// implements fmt.Stringer interface
func (gender gender) String() string {
	if gender == male {
		return "male"
	}

	return "female"
}
