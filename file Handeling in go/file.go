package main

import (
	"fmt"
	"os"
)

func main() {
	content := "Hello, Go!, Hello, Go! Hello, Go! Hello, Go! Hello, Go! Hello, Go! Hello, Go! "
	os.WriteFile("example.txt", []byte(content), 0644)
	data, err := os.ReadFile("example.txt") // Read the file
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(data)) // Print file content
}
