package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	fmt.Println("--- 6 LOOP ---")

	const age = 26
	for year := 0; year < age; year++ {
		fmt.Println(fmt.Sprintf("Your %d year of life", year+1))
	}

	actual, max := 26, 32

	for ; actual < max; actual += 1 {
		fmt.Println(actual)
	}

	for actual < max {
		actual += 1
		fmt.Println(actual)
	}

	value := sqrt(-12.10)

	fmt.Println(fmt.Sprintf("Value is %g", value))
	/*
	   	for {
	         fmt.Println("FOREVER LOOP KILL ME")
	   	}
	*/

	fmt.Println(pow(2, 14, 1000))

	personName, personAge := login(2)
	fmt.Printf("Welcome! %s %d years old", personName, personAge)
	fmt.Println()

	fmt.Println(fmt.Sprintf("%s: %s", "A", "Ahoj!"))
	fmt.Println(fmt.Sprintf("%s: %s", "B", hello(time.Now())))

	for i := 0; i < 3; i += 1 {
		fmt.Println()
	}

	tryDefer()

	for i := 0; i < 3; i += 1 {
		fmt.Println()
	}

	lifo()

}

func sqrt(x float64) float64 {
	if x > 0 {
		return math.Sqrt(x)
	} else {
		return -1
	}
}

func pow(x, n, lim float64) (value float64) {
	if value = math.Pow(x, n); value < lim {
		return value
	}
	return lim
}

func login(id int) (name string, age int) {
	switch id {
	case 1:
		name = "Ales"
		age = 26
	case 2:
		name = "Alena"
		age = 27
	case 3:
		name = "Petr"
		age = 36
	default:
		name = "Not Found"
		age = -1
	}
	return
}

func tryDefer() {
	defer fmt.Println("Finished!")
	fmt.Println("Prvni")
	fmt.Println("Druhy")
	fmt.Println("Treti")
}

func lifo() {
	fmt.Println("Start")
	for i := 0; i < 10; {
		i += 1
		defer fmt.Println(fmt.Sprintf("I'm %d!", i))
	}
	fmt.Println("Finished")
}

func hello(time time.Time) (answer string) {
	switch {
	case time.Hour() < 12:
		answer = "Dobre dopoledne"
	case time.Hour() == 12:
		answer = "Dobre poledne"
	case time.Hour() > 12:
		answer = "Dobre odpoledne"
	}
	return
}
