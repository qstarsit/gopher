package main

import (
	api "github.com/qstarsit/client"
)

func main() {
	client := api.Client{ServerAddress: "wss://gopher.arcetta.dev"} // wss is the Websocket equivalent of https

	// TODO: Connect to the server via the client, save the connection to a variable and use it instead of PARAMETER in StartReading

	/* This channel is used to block main from exiting before reading is done
	Without it, our main function will exit which causes our program to stop before all goroutines are finished.
	*/
	done := make(chan struct{}) 
	client.StartReading(PARAMETER, func(msg string) {
		// Implement a callback to handle incoming server messages
	})

	// Send messages to the websocket server from main thread via the client

	/* This will block main from exiting so we ensure the StartReading goroutine continues running in the background.
	Make sure to CTRL + C after the server sent back all your messages or you will be waiting for a while :-)
	*/
	<-done
}
