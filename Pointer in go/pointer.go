package main

import "fmt"

func main() {
	x := 1
	p := &x
	fmt.Println("Value of x:", *p)
	fmt.Println("Address of x:", &x)
}
