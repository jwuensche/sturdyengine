package sturdyengine

import (
	"flag"
	"fmt"
	"net/url"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
)

func (conn *Connection) InitializeAPI(clientName string) (e error) {
	conn.Upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	q := url.Values{}
	q.Add("name", clientName)

	conn.Uri = url.URL{
		Scheme: "ws", Host: *flag.String("addr", "localhost:50000", "kRPC client"),
		RawQuery: q.Encode(),
	}

	c, _, e := websocket.DefaultDialer.Dial(conn.Uri.String(), nil)
	conn.Conn = c

	return
}

func (conn *Connection) Close() (e error) {

	e = conn.Conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if e != nil {
		return
	}

	return
}

//GetStatus calls procedure `GetStatus` and returns the response in form of the proto buffer message
func (conn *Connection) GetStatus() (res Response, e error) {

	pc := ProcedureCall{
		Service:   "KRPC",
		Procedure: "GetStatus",
	}
	pr := Request{
		Calls: []*ProcedureCall{&pc},
	}

	res, e = conn.sendMessage(pr)

	fmt.Println(res.String())

	e = nil
	return
}

func (conn *Connection) sendMessage(r Request) (res Response, e error) {
	req, e := proto.Marshal(&r)

	if e != nil {
		return
	}

	conn.Conn.WriteMessage(websocket.BinaryMessage, req)

	_, p, e := conn.Conn.ReadMessage()
	if e != nil {
		return
	}
	res = Response{}
	e = proto.Unmarshal(p, &res)
	if e != nil {
		return
	}

	return
}
