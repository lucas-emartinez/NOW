package sse

import (
	"NOW/logic/entities/sse"
	"time"
)

type SSERepositoryImplementation struct {
	sse *sse.SSE
}

func NewSSEImplementation(sse *sse.SSE) *SSERepositoryImplementation {
	return &SSERepositoryImplementation{sse: sse}
}

func (s *SSERepositoryImplementation) AddClient(clientChan string) {

}

func (s *SSERepositoryImplementation) RemoveClient(clientChan string) {
	panic("implement me")
}

func (s *SSERepositoryImplementation) Broadcast(message string) {
	panic("implement me")
}

func (s *SSERepositoryImplementation) Ping(clientChan string) {
	panic("implement me ")
	// implement a ping method that sends a ping message to the clientChan
	// if the clientChan is not responding, remove it from the sse.Clients map
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// send a ping message to the clientChan

		}
	}
}
