package sturdyengine

import "github.com/golang/protobuf/proto"

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
