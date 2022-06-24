package main

import (
	"fmt"
	"time"
)

func message(c chan string, s string) {
	for {
		time.Sleep(1 * time.Second)
		c <- s
	}
}

func main() {
	chan1 := make(chan string)
	chan2 := make(chan string)

	go message(chan1, "hi")
	go message(chan2, ":)")

	for {
		select {
		case msg1 := <-chan1:
			fmt.Println(msg1)
		case msg2 := <-chan2:
			fmt.Println(msg2)
		default:
			fmt.Println("Default")
		}
	}
}
