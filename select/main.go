package main

import (
	"fmt"
	"select/database"
	"time"
)

// func message(c chan string, s string) {
// 	for {
// 		time.Sleep(1 * time.Second)
// 		c <- s
// 	}
// }

func main() {
	channel := make(chan bool)

	go database.Save("hi :)", channel)

	for {
		select {
		case <-channel:
			fmt.Println("Saved")
			return
		case <-time.After(3 * time.Second):
			fmt.Println("This is taking too long :(")
			return
		}
	}

	// chan1 := make(chan string)
	// chan2 := make(chan string)

	// go message(chan1, "hi")
	// go message(chan2, ":)")

	// for {
	// 	select {
	// 	case msg1 := <-chan1:
	// 		fmt.Println(msg1)
	// 	case msg2 := <-chan2:
	// 		fmt.Println(msg2)
	// 	default:
	// 		fmt.Println("Default")
	// 	}
	// }
}
