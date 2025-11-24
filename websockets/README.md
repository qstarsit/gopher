# Go WebSockets

In this exercise we will connect to a [WebSocket](https://ably.com/topic/websockets) server and exchange messages. A WebSocket begins as a normal HTTP request, you have to include the Upgrade header so the server knows to switch protocols and provide a full‑duplex channel (bidirectional communication) over the existing TCP connection.

## The Server

The server will accept your request and open a socket with you for `10` seconds. In this time, you can send messages to the server which will echo back all of your sent messages before closing the connection.

## API Client

We will make use of a Websocket API client you can find on [GitHub](https://github.com/qstarsit/client@latest). Make sure you have the API client included in your project with `go get github.com/qstarsit/client@latest`. 

This project already has it added for you, check out the `go.mod` file. To simply download the dependency as it's already declared, run `go mod download`.

## Task

There's a Websocket server running at `gopher.arcetta.dev`. You need to connect to it via the API Client. Inside `main.go` you will find the imported module, now it's up to you to connect to the server and send/read messages.

How to do so? Either read the original source code in the Github project above or find out what functions and methods are available via [pkg.go.dev](https://pkg.go.dev/github.com/qstarsit/client)

## Extra Challenge

The server also has a `/health` endpoint (http://gopher.arcetta.dev/health). Can you write a normal HTTP connecton with [`net/http`](https://pkg.go.dev/net/http) and get the health status of the app?

Even better, can you create a custom struct that holds the `health` endpoint data? Use `json.Marshal` ([tutorial](https://gobyexample.com/json)) to populate your struct with the response object.

## Compile time

After finishing the above tasks you now have a working go application. It's time to [compile](https://go.dev/doc/tutorial/compile-install) your app to run it natively on your machine.
