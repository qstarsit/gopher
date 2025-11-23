# Go WebSockets

In this exercise we will connect to a [WebSocket](https://ably.com/topic/websockets) server and exchange messages. A WebSocket begins as a normal HTTP request, you have to include the Upgrade header so the server knows to switch protocols and provide a full‑duplex channel (bidirectional communication) over the existing TCP connection.

## API Client

We will make use of a Websocket API client you can find on [GitHub](https://github.com/qstarsit/client@latest). Make sure you have the API client included in your project with `go get github.com/qstarsit/client@latest`.
