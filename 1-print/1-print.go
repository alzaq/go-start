// CMD+SHIFT+P -> Run File
package main

import (
	"fmt" // https://golang.org/pkg/fmt/
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("--- 1 PRINT ---")
	fmt.Printf("hello, world\n")
	fmt.Println("hello", "world", "How are you?", time.Now().Day())
	fmt.Printf("How much is %g Bitcon to USD?", rand.Float32())
	fmt.Println()

	fmt.Print(math.E)
	fmt.Println()

	rand.Seed(342824)
	fmt.Print(rand.ExpFloat64())
	fmt.Println()
}
