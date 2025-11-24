package examples

import (
	"fmt"

	api "github.com/qstarsit/client"
)

func main() {
	client := api.Client{ServerAddress: "wss://gopher.arcetta.dev"} // wss is the Websocket equivalent of https

	conn := client.Connect()
	defer conn.Close()

	// Start reading messages in background
	done := make(chan struct{}) // This channel is used to block main from exiting
	client.StartReading(conn, func(msg string) {
		fmt.Printf("Received: %s\n", msg)
	})

	// Send messages from main thread
	client.SendMessage(conn, "Test message #1")
	client.SendMessage(conn, "Test message #2")

	<-done // Block main from exiting until done is closed
}
