package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net"
	"net/http"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

type WebsocketMessage struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

type Message struct {
	Type   string `json:"type"`
	Data   string `json:"data"`
	Sender *websocket.Conn
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if err != nil {
		log.Println("Error during WebSocket upgrade:", err)
		return
	}
	defer func() {
		ws.Close()
		delete(clients, ws)
	}()

	clients[ws] = true

	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			delete(clients, ws)
			break
		}
		msg.Sender = ws
		broadcast <- msg
	}
}

func handleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			if client != msg.Sender {
				err := client.WriteJSON(msg)
				if err != nil {
					client.Close()
					delete(clients, client)
				}
			}
		}
	}
}

func handleIncomingData() {
	listener, err := net.Listen("tcp", "127.0.0.1:9000")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go readData(conn)
	}
}

func readData(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		return
	}

	var messageMap map[string]interface{}
	err = json.Unmarshal(buf[:n], &messageMap)
	if err != nil {
		log.Println("Error al deserializar el mensaje:", err)
		return
	}

	log.Println("Mensaje recibido:", messageMap)

	// Si necesitas convertir el mapa a una cadena JSON para enviarlo:
	messageJSON, err := json.Marshal(messageMap)
	if err != nil {
		log.Println("Error al serializar el mensaje:", err)
		return
	}

	broadcast <- Message{Data: string(messageJSON)}
}

func main() {
	http.HandleFunc("/ws", handleConnections)
	go handleMessages()
	go handleIncomingData()

	log.Println("WebSocket server started on :8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
