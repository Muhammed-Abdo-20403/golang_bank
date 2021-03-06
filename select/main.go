package main

import (
	"fmt"
	"time"
)

// func message(c chan string, s string) {
// 	for {
// 		time.Sleep(1 * time.Second)
// 		c <- s
// 	}
// }

// func message(c chan string, s string) {
// 	for {
// 		time.Sleep(500 * time.Millisecond)
// 		c <- s
// 	}
// }

func echo(done <-chan bool, interval time.Duration) {
	ticker := time.NewTicker(interval)
	for {
		select {
		case t := <-ticker.C:
			fmt.Println("Ticker tick ", t.Second())
			// get data from API
		case <-done:
			ticker.Stop()
			return
		}
	}
}

func main() {
	// channel := make(chan bool)

	// go database.Save("hi :)", channel)

	// for {
	// 	select {
	// 	case <-channel:
	// 		fmt.Println("Saved")
	// 		return
	// 	case <-time.After(3 * time.Second):
	// 		fmt.Println("This is taking too long :(")
	// 		return
	// 	}
	// }

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

	// chan1 := make(chan string)
	// chan2 := make(chan string)

	// go message(chan1, "hi")
	// go message(chan2, ":)")

	// timer := time.NewTimer(2 * time.Second)
	// for {
	// 	select {
	// 	case msg1 := <-chan1:
	// 		fmt.Println(msg1)
	// 	case msg2 := <-chan2:
	// 		fmt.Println(msg2)
	// 	case <-timer.C:
	// 		fmt.Println("Timer ended")
	// 		return
	// 	}
	// }

	// timer := time.NewTimer(2 * time.Second)
	// <-timer.C
	// fmt.Println("Timer kicked")

	//  you can cancel the timer before it kick
	// timer := time.NewTimer(time.Second)
	// go func() {
	// 	<-timer.C
	// 	fmt.Println("Timer kicked")
	// }()
	// stoped := timer.Stop()
	// if stoped {
	// 	fmt.Println("Timer stopped")
	// }
	// time.Sleep(3 * time.Second)

	// ticker := time.NewTicker(1 * time.Second)
	// timer := time.NewTimer(5 * time.Second)

	// for {
	// 	select {
	// 	case t := <-ticker.C:
	// 		fmt.Println("Ticker tick", t.Second())
	// 	case <-timer.C:
	// 		ticker.Stop()
	// 		return
	// 	}

	// }
	done := make(chan bool)
	go echo(done, time.Second*1)

	time.Sleep(5 * time.Second)
	done <- true
	fmt.Println("Ticker stopped")
}
