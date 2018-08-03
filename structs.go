package sturdyengine

import (
	"net/url"

	"github.com/gorilla/websocket"
)

type Connection struct {
	Uri url.URL

	Upgrader websocket.Upgrader
	Conn     *websocket.Conn
}
