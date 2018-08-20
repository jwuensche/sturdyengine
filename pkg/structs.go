package sturdyengine

import (
	"net/url"

	"github.com/gorilla/websocket"
)

// Connection contains all inforamtion required by websocket to establish connection
type Connection struct {
	URI url.URL

	Upgrader websocket.Upgrader
	Conn     *websocket.Conn
}

// SpaceCenter saves all information used by `SpaceCenter` service of kRPC
type SpaceCenter struct {
	conn    *Connection
	Vessel  []byte
	Control []byte
}
