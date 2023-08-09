package sse

type SSE struct {
	Clients map[chan string]bool
}

func NewSSEHandler() *SSE {
	return &SSE{
		Clients: make(map[chan string]bool),
	}
}
