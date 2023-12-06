package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a channel of strings
	c := make(chan string)

	// Create two goroutines
	go ping(c)
	go pong(c)

	// Wait for the goroutines to finish
	for {
		fmt.Println(<-c)
	}
}

func ping(c chan string) {
	for {
		c <- "ping"
		time.Sleep(1 * time.Second)
	}
}

func pong(c chan string) {
	for {
		c <- "pong"
		time.Sleep(1 * time.Second)
	}
}
