package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Define flags
	add := flag.Bool("add", false, "Perform addition")
	sub := flag.Bool("sub", false, "Perform subtraction")
	mul := flag.Bool("mul", false, "Perform multiplication")
	div := flag.Bool("div", false, "Perform division")
	flag.Parse()

	// Get remaining command-line arguments
	args := flag.Args()

	if len(args) != 2 {
		fmt.Println("Error: Please provide exactly two numbers")
		os.Exit(1)
	}

	// Convert arguments to numbers
	num1, err1 := strconv.ParseFloat(args[0], 64)
	num2, err2 := strconv.ParseFloat(args[1], 64)

	if err1 != nil || err2 != nil {
		fmt.Println("Error: Both inputs must be valid numbers")
		os.Exit(1)
	}

	// Perform the selected operation
	var result float64
	var err error

	if *add {
		result = num1 + num2
	} else if *sub {
		result = num1 - num2
	} else if *mul {
		result = num1 * num2
	} else if *div {
		if num2 == 0 {
			err = fmt.Errorf("Error: Division by zero")
		} else {
			result = num1 / num2
		}
	} else {
		err = fmt.Errorf("Error: Please specify an operation using flags (e.g., -add, -sub, -mul, -div)")
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Result: %.2f\n", result)
}
