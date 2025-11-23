package examples

import (
	"fmt"
	"time"
)

func goroutinesAndChannelsExample() {
	/*
		A channel is created to communicate between goroutines.
		It is a typed "conduit" through which you can send and receive values.
	*/
	messages := make(chan string)

	/*
		We start a new goroutine that sends a message into the channel.
	*/
	go func() {
		time.Sleep(time.Second) // Do some fake work by sleeping a second
		messages <- "Hello from the goroutine!"
	}()

	/*
		Main goroutine waits to receive a message from the channel.
		This receive operation blocks until a message is available.
	*/
	msg := <-messages
	fmt.Println("Received message:", msg) // Only printed after the channel was populated

	/*
		We can also close a channel when no more values will be sent.
		This signals receivers that no more data is coming.
	*/
	close(messages)
}
