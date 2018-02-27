package main

import "fmt"

var male bool
var name string = "Default"
var age int
var percentage float32

var (
	street        string
	zip           int
	czechRepublic bool
)

func main() {
	fmt.Println("--- 3 TYPES ---")

	male = false
	name = "Ales"
	age = 26
	percentage = 0.4920

	fmt.Println(name, male, age, (percentage * 100))

	street = "Hlavnov"
	zip = 54954
	czechRepublic = false

	fmt.Println("My Address", street, zip, czechRepublic)

	i := 27
	f := 2.8929
	u := uint(9310)

	fmt.Println(i, f, u)
	fmt.Println(i, float64(f), uint(float64(f)))

	var fl float64

	fl = float64(i)
	fmt.Println(fl)

	var (
		eI int
		eF float32
		eB bool
		eS string = "ahoj"
	)

	fmt.Println(eI, eF, eB, eS, len(eS))

}
