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

}
