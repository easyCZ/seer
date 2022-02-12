package api

import (
	"github.com/gorilla/websocket"
	"log"
	"os"
)

type clientPool struct {
	clients map[*client]bool

	register   chan *client
	unregister chan *client
}

func (p *clientPool) run() {
	logger := log.New(os.Stderr, "client_pool", log.LstdFlags)

	for {
		select {
		case c := <-p.register:
			p.clients[c] = true
			logger.Println("Client added. Now has", len(p.clients), "clients.")
		case c := <-p.unregister:
			delete(p.clients, c)
			logger.Println("Client removed. Now has", len(p.clients), "clients.")

			// TODO: close the client struct gracefully
		}
	}
}

type client struct {
	conn *websocket.Conn
}

type message struct {
	payload []byte
}
