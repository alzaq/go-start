package main

import (
	"fmt"
	"time"
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

func formateTime(date time.Time) string {
	return fmt.Sprintf("%d:%d:%d", date.Hour(), date.Minute(), date.Second())
}

func getNameOfDay(date time.Time) string {
	switch date.Weekday() {
	case 1:
		return "Pondeli"
	case 2:
		return "Utery"
	case 3:
		return "Streda"
	default:
		return fmt.Sprintf("Day: %d", date.Weekday())
	}
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

	fmt.Println(formateTime(time.Now()))

	var dayOfWeek = getNameOfDay(time.Now())
	fmt.Println(dayOfWeek)
}
