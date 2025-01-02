package main

import "fmt"

func greet(name string) string {
	return "You are great " + name
}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	output := greet("Sameer")
	greet2 := func(name string) {
		fmt.Println("Hello,", name)
	}
	greet2("Alice")
	fmt.Println(output)
	fmt.Println(sum(1, 2, 3, 4, 5))
}
