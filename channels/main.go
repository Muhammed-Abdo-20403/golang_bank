package main

import (
	"fmt"
	"sync"
	"time"
)

func mss(messages chan string) {
	time.Sleep(2 * time.Second)
	messages <- "Hi :)"
}

func main() {
	messages := make(chan string)

	go mss(messages)

	msg := <-messages
	fmt.Println(msg)

	// /***** rang & close Channels *****/
	messages = make(chan string, 2)
	messages <- "Hi"
	messages <- ":)"
	//close(messages)
	go func() {
		time.Sleep(time.Second * 2)
		close(messages)
	}()
	fmt.Println("execute without anyone reading the message yet")
	for message := range messages {
		fmt.Println(message)
	}
	fmt.Println("after range")

	var wg sync.WaitGroup
	ping := make(chan bool)
	pong := make(chan bool)

	go playPing(ping, pong, &wg)
	go playPong(ping, pong, &wg)

	// this will kik start one of the playPing()
	pong <- true
	wg.Wait()
}

func playPing(pingChan chan<- bool, pongChan <-chan bool, wg *sync.WaitGroup) {
	for {
		wg.Add(1)
		defer wg.Done()

		time.Sleep(500 * time.Millisecond)
		// wait to read on pongChan
		<-pongChan
		fmt.Println("Ping")
		pingChan <- true
	}
}
func playPong(pingChan <-chan bool, pongChan chan<- bool, wg *sync.WaitGroup) {
	for {
		wg.Add(1)
		defer wg.Done()

		time.Sleep(500 * time.Millisecond)
		<-pingChan
		fmt.Println("Pong")
		pongChan <- true
	}
}
