package sturdyengine

import (
	"fmt"

	"github.com/golang/protobuf/proto"
)

func (conn *Connection) GetGameMode() (r, e error) {
	pc := &ProcedureCall{
		Service:   "SpaceCenter.Vessel.Control",
		Procedure: "get_GameMode",
	}
	pr := &Request{
		Calls: []*ProcedureCall{pc},
	}

	p, e := conn.sendMessage(pr)
	res := &Response{}
	proto.Unmarshal(p, res)

	return
}

func (conn *Connection) GetActiveVessel() (r []byte, e error) {
	pc := &ProcedureCall{
		Service:   "SpaceCenter",
		Procedure: "get_ActiveVessel",
	}
	pr := &Request{
		Calls: []*ProcedureCall{pc},
	}

	p, e := conn.sendMessage(pr)
	res := &Response{}
	proto.Unmarshal(p, res)
	r = res.GetResults()[0].GetValue()
	return
}

func (conn *Connection) GetVesselControl(vessel []byte) (r []byte, e error) {
	pc := &ProcedureCall{
		Service:   "SpaceCenter",
		Procedure: "Vessel_get_Control",
		Arguments: []*Argument{&Argument{
			Position: 0,
			Value:    vessel,
		}},
	}
	pr := &Request{
		Calls: []*ProcedureCall{pc},
	}

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
	pc := &ProcedureCall{
		Service:   "SpaceCenter",
		Procedure: "set_SASMode",
	}
	pr := &Request{
		Calls: []*ProcedureCall{pc},
	}

	p, e := conn.sendMessage(pr)
	res := &Response{}
	proto.Unmarshal(p, res)
	fmt.Println(res.String())

	return
}

func (conn *Connection) ActivateNextStage(vessel []byte) (r, e error) {
	pc := &ProcedureCall{
		Service:   "SpaceCenter",
		Procedure: "Control_ActivateNextStage",
		Arguments: []*Argument{&Argument{
			Position: 0,
			Value:    vessel,
		}},
	}
	pr := &Request{
		Calls: []*ProcedureCall{pc},
	}

	p, e := conn.sendMessage(pr)
	if e != nil {
		return
	}
	res := &Response{}
	proto.Unmarshal(p, res)

	return
}
