package sturdyengine

import (
	"flag"
	"net/url"

	"github.com/gorilla/websocket"
)

func (conn *Connection) InitializeAPI(clientName string) (e error) {
	conn.Upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	// TODO: This is silly....
	u := url.URL{}
	q := u.Query()
	q.Set("name", clientName)

	conn.Uri = url.URL{
		Scheme: "ws", Host: *flag.String("addr", "localhost:50000", "kRPC client"),
		RawQuery: q.Encode(),
	}

	c, _, e := websocket.DefaultDialer.Dial(conn.Uri.String(), nil)
	conn.Conn = c

	return
}
