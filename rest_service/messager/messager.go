package messager

import (
	"encoding/json"
	"log"
	"net"
)

type WebsocketMessage struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

func SendMessageToWebSocket(message WebsocketMessage) error {
	conn, err := net.Dial("tcp", "localhost:9000") // Asume que el servidor WebSocket escucha en el puerto 9000
	if err != nil {
		log.Println("Error al conectar con el servidor WebSocket:", err)
		return err
	}
	defer conn.Close()

	// Serializa el mensaje a JSON
	jsonMessage, err := json.Marshal(message)
	if err != nil {
		log.Println("Error al serializar el mensaje:", err)
		return err
	}

	_, err = conn.Write(jsonMessage)
	if err != nil {
		log.Println("Error al enviar mensaje al servidor WebSocket:", err)
	}
	return err
}
