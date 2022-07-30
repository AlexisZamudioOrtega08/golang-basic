package main

import (
	"fmt"
	"time"
)

func main() {
	// a channel is a communication mechanism between goroutines.
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			// send a message to the channel
			c1 <- "routine 1"
			time.Sleep(time.Second * 1)
		}
	}()

	go func() {
		for {
			// send a message to the channel
			c2 <- "routine 2"
			time.Sleep(time.Second * 1)
		}
	}()

	for x := 0; x < 10; x++ {
		// select statement is used to wait for a channel to be ready for sending or receiving and not block the program.
		select {
		// <-c1 is used to receive a message from the channel.
		case msg1 := <-c1:
			fmt.Println(msg1)
		// <-c1 is used to receive a message from the channel.
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}

}
