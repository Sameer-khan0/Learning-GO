package mypackage

import "fmt"

// Hello prints a greeting message.
func Hello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// Add adds two integers and returns the result.
func Add(a, b int) int {
	return a + b
}
