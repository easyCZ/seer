package agent

import "net/url"

func RunClient() {
	u := url.URL{Scheme: "ws", Host: "http://localhost:3000/", Path: "/echo"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
}
