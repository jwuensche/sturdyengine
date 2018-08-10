package sturdyengine

import (
	"fmt"

	"github.com/golang/protobuf/proto"
)

func (conn *Connection) GetGameMode() (r, e error) {
	pr := createRequest("SpaceCenter", "get_GameMode", nil)

	p, e := conn.sendMessage(pr)
	res := &Response{}
	proto.Unmarshal(p, res)

	return
}

func (conn *Connection) GetActiveVessel() (r []byte, e error) {
	pr := createRequest("SpaceCenter", "get_ActiveVessel", nil)

	p, e := conn.sendMessage(pr)
	res := &Response{}
	proto.Unmarshal(p, res)
	r = res.GetResults()[0].GetValue()
	return
}

func (conn *Connection) GetVesselControl(vessel []byte) (r []byte, e error) {
	arg := []*Argument{&Argument{
		Position: 0,
		Value:    vessel,
	}}
	pr := createRequest("SpaceCenter", "Vessel_get_Control", arg)

	p, e := conn.sendMessage(pr)
	if e != nil {
		return
	}
	res := &Response{}
	e = proto.Unmarshal(p, res)
	if e != nil {
		return
	}
	r = res.GetResults()[0].GetValue()
	return
}

func (conn *Connection) SetSAS() (r, e error) {
	pr := createRequest("SpaceCenter", "set_SASMode", nil)

	p, e := conn.sendMessage(pr)
	res := &Response{}
	proto.Unmarshal(p, res)
	fmt.Println(res.String())

	return
}

func (conn *Connection) ActivateNextStage(vessel []byte) (e error) {

	arg := []*Argument{&Argument{
		Position: 0,
		Value:    vessel,
	}}
	pr := createRequest("SpaceCenter", "Control_ActivateNextStage", arg)

	p, e := conn.sendMessage(pr)
	if e != nil {
		return
	}
	res := &Response{}
	proto.Unmarshal(p, res)

	return
}

func createRequest(service string, procedure string, arguments []*Argument) (pr *Request) {
	pc := &ProcedureCall{
		Service:   service,
		Procedure: procedure,
		Arguments: arguments,
	}
	pr = &Request{
		Calls: []*ProcedureCall{pc},
	}
	return
}