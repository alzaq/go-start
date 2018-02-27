package main

import (
	"fmt"
)

func genName(char string, number int) string {
	var name = ""
	for index := 0; index < number; index++ {
		name += char
	}
	return name
}

func minMax(x, y int) (int, int) {
	if x > y {
		return y, x
	} else {
		return x, y
	}
}

func valInfo(x, y int) (min int, max int, avg int) {
	min, max = minMax(x, y)
	avg = (min + max) / 2
	return
}

func main() {
	fmt.Println("--- 2 FUNC ---")
	fmt.Printf("Hello %s!", "Ales")
	fmt.Println()
	fmt.Println(genName("A", 12))

	a := 26
	b := 24

	fmt.Println(minMax(a, b))

	min, max, avg := valInfo(a, b)
	fmt.Println(min, max, avg)
}
