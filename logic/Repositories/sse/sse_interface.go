package sse

type SSERepository interface {
	AddClient(clientChan chan string)
	RemoveClient(clientChan chan string)
	Ping(clientChan chan string)
	Broadcast(message string)
}
