package sturdyengine

import (
	"github.com/golang/protobuf/proto"
	krpc "github.com/jwuensche/sturdyengine/internal/krpcproto"
)

//GetStatus calls remote procedure `GetStatus` and returns the response in form of the proto buffer message
func (conn *Connection) GetStatus() (res krpc.Status, e error) {
	pc := krpc.ProcedureCall{
		Service:   "KRPC",
		Procedure: "GetStatus",
	}
	pr := krpc.Request{
		Calls: []*krpc.ProcedureCall{&pc},
	}

	p, e := conn.sendMessage(&pr)
	if e != nil {
		return
	}

	r := krpc.Response{}
	e = proto.Unmarshal(p, &r)
	if e != nil {
		return
	}
	e = proto.Unmarshal(r.GetResults()[0].GetValue(), &res)

	return
}

//GetServices calls KRPC remote procedure `GetServices` and returns Service Structure
func (conn *Connection) GetServices() (res krpc.Services, e error) {
	pc := krpc.ProcedureCall{
		Service:   "KRPC",
		Procedure: "GetServices",
	}
	pr := krpc.Request{
		Calls: []*krpc.ProcedureCall{&pc},
	}

	p, e := conn.sendMessage(&pr)
	if e != nil {
		return
	}

	r := &krpc.Response{}
	e = proto.Unmarshal(p, r)
	if e != nil {
		return
	}
	e = proto.Unmarshal(r.GetResults()[0].GetValue(), &res)
	return
}
