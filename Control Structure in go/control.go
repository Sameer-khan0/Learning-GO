package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	for index, value := range nums {
		fmt.Printf("index: %d, value: %d\n", index, value)
	}
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
