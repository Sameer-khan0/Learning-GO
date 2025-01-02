package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	jsonData := `{"Name":"Alice","Age":25}`
	var p Person
	json.Unmarshal([]byte(jsonData), &p)
	fmt.Println(p.Name, "is", p.Age, "years old.")
}
