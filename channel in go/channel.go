package main

import "fmt"

func sendMessage(ch chan string) {
	ch <- "Hello from goroutine!" // Send data to channel
}

func main() {
	ch := make(chan string) // Create a channel
	go sendMessage(ch)      // Start a goroutine
	msg := <-ch             // Receive data from channel
	fmt.Println(msg)
}
