package main

import "fmt"

type Animal interface {
	Says() string
	NumberOfLegs() int
}
type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {
	dog := Dog{
		Name:  "hello",
		Breed: "Dhruv Rathee",
	}
	PrintInfo(&dog)
	gorilla := Gorilla{
		Name:          "cha",
		Color:         "grey",
		NumberOfTeeth: 38,
	}
	PrintInfo(&gorilla)
}

func PrintInfo(a Animal) {
	fmt.Println("this animal says", a.Says(), "and has", a.NumberOfLegs(), "legs")
}
func (d *Dog) Says() string {
	return "woof"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}
func (d *Gorilla) Says() string {
	return "ugh"
}

func (d *Gorilla) NumberOfLegs() int {
	return 2
}
