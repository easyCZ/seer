package agent

import (
	"fmt"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

func RunClient() error {
	u := url.URL{Scheme: "ws", Host: "localhost:3000", Path: "/ws/agents"}
	log.Printf("connecting to %s", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return fmt.Errorf("failed to establish websocket connection: %w", err)
	}

	if err := conn.WriteMessage(websocket.TextMessage, []byte(`yolo`)); err != nil {
		log.Println("write:", err)
		return err
	}

	if err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")); err != nil {
		log.Println("write close:", err)
		return err
	}

	return nil
}
