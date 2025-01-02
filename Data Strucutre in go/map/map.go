package main

import "fmt"

func main() {
	capital := map[string]string{
		"pakistan": "Islamabad",
		"USA":      "Washington D.C.",
		"India":    "New Delhi",
		"Japan":    "Tokyo",
	}
	capital["France"] = "Paris"
	capital["India"] = "Delhi"
	fmt.Println(capital["India"])
	fmt.Println(len(capital))
	fmt.Println(capital)
}
