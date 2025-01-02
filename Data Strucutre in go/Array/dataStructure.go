package main

import "fmt"

func main() {
	// var str [5]string = [5]string{"A", "B", "Cwwwwww", "Dfwefwefewfew", "E"}
	num := []int{1, 2, 3, 4, 5}
	num = append(num, 10)
	num = append(num, 10)
	num = append(num, 10)
	fmt.Println(num[0:3])
}
