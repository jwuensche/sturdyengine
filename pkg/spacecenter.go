package sturdyengine

import (
	"github.com/golang/protobuf/proto"
)

func InitSpaceCenter(conn *Connection) (sc SpaceCenter, e error) {
	sc = SpaceCenter{}
	sc.conn = conn
	sc.Vessel, e = sc.GetActiveVessel()
	sc.Control, e = sc.GetVesselControl(sc.Vessel)
	return
}

func (sc *SpaceCenter) GetGameMode() (r []byte, e error) {
	pr := createRequest("SpaceCenter", "get_GameMode", nil)

	p, e := sc.conn.sendMessage(pr)
	res := &Response{}
	proto.Unmarshal(p, res)

	return
}

func (sc *SpaceCenter) GetActiveVessel() (r []byte, e error) {
	pr := createRequest("SpaceCenter", "get_ActiveVessel", nil)

	p, e := sc.conn.sendMessage(pr)
	res := &Response{}
	proto.Unmarshal(p, res)
	r = res.GetResults()[0].GetValue()
	return
}

func (sc *SpaceCenter) GetVesselControl(vessel []byte) (r []byte, e error) {
	arg := []*Argument{&Argument{
		Position: 0,
		Value:    vessel,
	}}
	pr := createRequest("SpaceCenter", "Vessel_get_Control", arg)

	p, e := sc.conn.sendMessage(pr)
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

// CONTROL - SPACECENTER.VESSEL.CONTROL

func (sc *SpaceCenter) SetSAS(vessel []byte, state bool) (r, e error) {
	var s []byte
	if state {
		s = []byte{byte(1)}
	} else {
		s = []byte{byte(0)}
	}
	arg := []*Argument{
		&Argument{
			Position: 0,
			Value:    vessel,
		},
		&Argument{
			Position: 1,
			Value:    s,
		},
	}
	pr := createRequest("SpaceCenter", "Control_set_SAS", arg)

	p, e := sc.conn.sendMessage(pr)
	res := &Response{}
	proto.Unmarshal(p, res)

	return
}

func (sc *SpaceCenter) ActivateNextStage(vessel []byte) (e error) {

	arg := []*Argument{&Argument{
		Position: 0,
		Value:    vessel,
	}}
	pr := createRequest("SpaceCenter", "Control_ActivateNextStage", arg)

	p, e := sc.conn.sendMessage(pr)
	if e != nil {
		return
	}
	res := &Response{}
	proto.Unmarshal(p, res)

	return
}

func (sc *SpaceCenter) GetThrottle() (val float32, e error) {
	arg := []*Argument{&Argument{
		Position: 0,
		Value:    sc.Control,
	}}
	pr := createRequest("SpaceCenter", "Control_get_Throttle", arg)
	p, e := sc.conn.sendMessage(pr)
	if e != nil {
		return
	}
	res := &Response{}
	proto.Unmarshal(p, res)
	val = byteToFloat32(res.GetResults()[0].GetValue())
	return
}

func (sc *SpaceCenter) SetThrottle(val float32) (e error) {
	arg := []*Argument{
		&Argument{
			Position: 0,
			Value:    sc.Control,
		},
		&Argument{
			Position: 1,
			Value:    float32toByte(val),
		},
	}

	pr := createRequest("SpaceCenter", "Control_set_Throttle", arg)
	_, e = sc.conn.sendMessage(pr)

	return
}

func (sc *SpaceCenter) Quicksave() (e error) {
	pr := createRequest("SpaceCenter", "Quicksave", nil)
	_, e = sc.conn.sendMessage(pr)
	return
}

func (sc *SpaceCenter) Quickload() (e error) {
	pr := createRequest("SpaceCenter", "Quickload", nil)
	_, e = sc.conn.sendMessage(pr)
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
