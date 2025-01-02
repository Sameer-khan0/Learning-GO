package main

import (
	"fmt"
	"math"
)

func main() {
	var num int = 16
	var sqr float64 = math.Sqrt(float64(num)) // Converting int to float64
	fmt.Printf("The square root of %d is %.2f\n", num, sqr)
	// result := mypackage.Add(3, 4)
	// fmt.Println("Result:", result)
}
