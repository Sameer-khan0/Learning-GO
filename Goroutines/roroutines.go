package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep((500 * time.Millisecond))
	}
}
func main() {
	go printNumbers() // Start a goroutine
	fmt.Println("Main function running concurrently")
	time.Sleep(3 * time.Second) // Give time for the goroutine to complete
}
