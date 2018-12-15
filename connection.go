package sturdyengine

import (
	"errors"
	"flag"
	"net/url"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	krpc "github.com/jwuensche/sturdyengine/internal/krpcproto"
)

// NewConnection sets informations required by the connection and also creates said. If the connection cannot be established returns error and struct member Conn is unset
func NewConnection(clientName string, port string) (conn Connection, e error) {
	conn = Connection{}
	conn.Upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	q := url.Values{}
	q.Add("name", clientName)

	f := flag.NewFlagSet(clientName, flag.ExitOnError)

	conn.URI = url.URL{
		Scheme: "ws", Host: *f.String("addr", "localhost:"+port, "kRPC client"),
		RawQuery: q.Encode(),
	}

	conn.Conn, _, e = websocket.DefaultDialer.Dial(conn.URI.String(), nil)
	return
}

// NewDefaultConnection simply calls InitializeAPI with the default port of the krpc mod
func NewDefaultConnection() (conn Connection, e error) {
	conn, e = NewConnection("sturdyengine", "50000")
	return
}

// Close closes the current connection
func (conn *Connection) Close() (e error) {
	if conn.Conn == nil {
		e = errors.New("Connection is nil")
		return
	}
	e = conn.Conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if e != nil {
		return
	}

	return
}

func (conn *Connection) sendMessage(r *krpc.Request) (p []byte, e error) {
	req, e := proto.Marshal(r)
	if e != nil {
		return
	}
	if conn.Conn == nil {
		e = errors.New("Connection is nil, check if any server is running on given port")
		return
	}

	conn.Conn.WriteMessage(websocket.BinaryMessage, req)
	_, p, e = conn.Conn.ReadMessage()
	if e != nil {
		return
	}

	return
}
