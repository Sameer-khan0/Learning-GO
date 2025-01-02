package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	p := Person{name: "Sameer", age: 21}
	p2 := Person{name: "Ali", age: 22}
	fmt.Println(p.name, p.age)
	fmt.Println(p2.name, p2.age)
}
