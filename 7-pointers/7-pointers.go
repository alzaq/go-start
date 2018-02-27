package main

import "fmt"

func main() {
	fmt.Println("--- 7 POINTERS ---")

	var p *int

	age := 26
	p = &age

	age += 1
	fmt.Printf("Age is: %d. Pointer is: %d. Value in pointer is: %d", age, p, *p)
	fmt.Println()

	*p += 1
	fmt.Printf("Age is: %d. Pointer is: %d. Value in pointer is: %d", age, p, *p)
	fmt.Println()
}
