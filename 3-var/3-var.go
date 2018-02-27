package main

import "fmt"

var min, max int = 1, 100

func main() {
	fmt.Println("--- 3 VARS ---")

	fmt.Println(fmt.Sprintf("Min: %d and Max: %d", min, max))

	var male, female, age, name = true, false, 27, "Ales"

	var gender string

	if male && !female {
		gender = "boy"
	} else if female && !male {
		gender = "girl"
	} else {
		gender = "ladyboy"
	}

	address := "Hlavnov"

	fmt.Printf("My name is %s, I'm %d years old and I am %s", name, age, gender)
	fmt.Println()
	fmt.Println(fmt.Sprintf("I'm living in %s", address))

}
