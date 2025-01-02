package main

import "fmt"

func main() {
	var age int = 1
	grade := 'C'
	if age >= 10 {
		fmt.Println("You are young boy")
	} else {
		fmt.Println("You are child")
	}
	for i := 0; i < 10; i++ {
		fmt.Println("The itrative index is ", i)
	}
	switch grade {
	case 'A':
		fmt.Println("Very Good")
	case 'B':
		fmt.Println("Good")
	case 'C':
		fmt.Println("Bad")
	}

}
