package main

import "fmt"

func greet(name string) string {
	return "You are great " + name
}
func main() {
	output := greet("Sameer")
	fmt.Println(output)
}
