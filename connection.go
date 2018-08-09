package sturdyengine

import (
	"flag"
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

//GetStatus calls remote procedure `GetStatus` and returns the response in form of the proto buffer message
func (conn *Connection) GetStatus() (res Status, e error) {
	pc := ProcedureCall{
		Service:   "KRPC",
		Procedure: "GetStatus",
	}
	pr := Request{
		Calls: []*ProcedureCall{&pc},
	}

	p, e := conn.sendMessage(&pr)
	if e != nil {
		return
	}

	r := Response{}
	e = proto.Unmarshal(p, &r)
	if e != nil {
		return
	}
	e = proto.Unmarshal(r.GetResults()[0].GetValue(), &res)

	return
}

func (conn *Connection) GetServices() (res Services, e error) {
	pc := ProcedureCall{
		Service:   "KRPC",
		Procedure: "GetServices",
	}
	pr := Request{
		Calls: []*ProcedureCall{&pc},
	}

	p, e := conn.sendMessage(&pr)
	if e != nil {
		return
	}

	r := &Response{}
	e = proto.Unmarshal(p, r)
	if e != nil {
		return
	}
	e = proto.Unmarshal(r.GetResults()[0].GetValue(), &res)

	return
}

func (conn *Connection) sendMessage(r *Request) (p []byte, e error) {
	req, e := proto.Marshal(r)

	if e != nil {
		return
	}

	conn.Conn.WriteMessage(websocket.BinaryMessage, req)

	_, p, e = conn.Conn.ReadMessage()
	if e != nil {
		return
	}

	return
}
