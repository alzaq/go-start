package main

import "fmt"

const PI = 3.14
const TAG = "CONSTS"

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	BIG = 1 << 5
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	SMALL = BIG >> 4
)

func createMessage(typ, text string) string {
	return fmt.Sprintf("%s[%s] - %s", TAG, typ, text)
}

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) (value float64) {
	value = x * 0.1
	return
}

func main() {
	fmt.Println("--- 5 CONSTS ---")

	var msg string
	msg = fmt.Sprintf("PI is about %g!", PI)
	fmt.Println(msg)

	fmt.Println(fmt.Sprintf("%s - %s", TAG, "There was some problem!"))

	const (
		E = "ERROR"
		W = "WARNING"
		I = "INFO"
	)

	fmt.Println(createMessage(E, "OH NO! BUG IS HERE!"))
	fmt.Println(createMessage(W, "YOU MADE A MISTAKE"))
	fmt.Println(createMessage(I, "YOU CAN USE SOMETHING BETTER"))

	fmt.Println()
	fmt.Println(needInt(SMALL), needFloat(SMALL), needInt(BIG), needFloat(BIG))

}
