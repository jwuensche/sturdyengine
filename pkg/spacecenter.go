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

// ORBIT - SPACECENTER.VESSEL.ORBIT || SPACECENTER.CELESTIALBODY.ORBIT

func (sc *SpaceCenter) GetVesselOrbit(vessel []byte) (orb []byte, e error) {
	arg := [][]byte{vessel}
	pr := createRequest("SpaceCenter", "Vessel_get_Orbit", createArguments(arg))
	p, e := sc.conn.sendMessage(pr)
	res := &Response{}
	proto.Unmarshal(p, res)
	orb = res.GetResults()[0].GetValue()
	return
}

func (sc *SpaceCenter) GetApoapsisAltitude(orbit []byte) (alt float64, e error) {
	alt, e = sc.getOrbitInfo(orbit, "Orbit_get_ApoapsisAltitude")
	return
}

func (sc *SpaceCenter) GetPeriapsisAltitude(orbit []byte) (alt float64, e error) {
	alt, e = sc.getOrbitInfo(orbit, "Orbit_get_PeriapsisAltitude")
	return
}

func (sc *SpaceCenter) GetSemiMajorAxis(orbit []byte) (alt float64, e error) {
	alt, e = sc.getOrbitInfo(orbit, "Orbit_get_SemiMajorAxis")
	return
}

func (sc *SpaceCenter) GetSemiMinorAxis(orbit []byte) (alt float64, e error) {
	alt, e = sc.getOrbitInfo(orbit, "Orbit_get_SemiMinorAxis")
	return
}

func (sc *SpaceCenter) GetRadius(orbit []byte) (alt float64, e error) {
	alt, e = sc.getOrbitInfo(orbit, "Orbit_get_Radius")
	return
}

func (sc *SpaceCenter) GetSpeed(orbit []byte) (alt float64, e error) {
	alt, e = sc.getOrbitInfo(orbit, "Orbit_get_Speed")
	return
}

func (sc *SpaceCenter) GetPeriod(orbit []byte) (time float64, e error) {
	time, e = sc.getOrbitInfo(orbit, "Orbit_get_Period")
	return
}

func (sc *SpaceCenter) GetTimeToApoapsis(orbit []byte) (time float64, e error) {
	time, e = sc.getOrbitInfo(orbit, "Orbit_get_TimeToApoapsis")
	return
}

func (sc *SpaceCenter) GetTimeToPeriapsis(orbit []byte) (time float64, e error) {
	time, e = sc.getOrbitInfo(orbit, "Orbit_get_TimeToPeriapsis")
	return
}

func (sc *SpaceCenter) GetEpoch(orbit []byte) (time float64, e error) {
	time, e = sc.getOrbitInfo(orbit, "Orbit_get_Epoch")
	return
}

func (sc *SpaceCenter) GetEccentricty(orbit []byte) (ecc float32, e error) {
	ecc, e = sc.getOrbitInfoFloat32(orbit, "Orbit_get_Eccentricity")
	return
}

func (sc *SpaceCenter) getOrbitInfo(orbit []byte, procedure string) (alt float64, e error) {
	arg := [][]byte{orbit}
	pr := createRequest("SpaceCenter", procedure, createArguments(arg))
	p, e := sc.conn.sendMessage(pr)
	if e != nil {
		return
	}
	res := &Response{}
	proto.Unmarshal(p, res)
	alt = byteToFloat64(res.GetResults()[0].GetValue())
	return
}

func (sc *SpaceCenter) getOrbitInfoFloat32(orbit []byte, procedure string) (alt float32, e error) {
	arg := [][]byte{orbit}
	pr := createRequest("SpaceCenter", procedure, createArguments(arg))
	p, e := sc.conn.sendMessage(pr)
	if e != nil {
		return
	}
	res := &Response{}
	proto.Unmarshal(p, res)
	alt = byteToFloat32(res.GetResults()[0].GetValue())
	return
}

// CONTROL - SPACECENTER.VESSEL.CONTROL

func (sc *SpaceCenter) GetVesselControl(vessel []byte) (r []byte, e error) {
	arg := [][]byte{vessel}
	pr := createRequest("SpaceCenter", "Vessel_get_Control", createArguments(arg))

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

func (sc *SpaceCenter) GetSAS(control []byte) (r bool, e error) {
	arg := [][]byte{control}
	pr := createRequest("SpaceCenter", "Control_get_SAS", createArguments(arg))
	p, e := sc.conn.sendMessage(pr)
	res := &Response{}
	proto.Unmarshal(p, res)
	r = byteToBool(res.GetResults()[0].GetValue())
	return
}

func (sc *SpaceCenter) SetSAS(vessel []byte, state bool) (e error) {
	s := boolToByte(state)
	arg := [][]byte{vessel, s}
	pr := createRequest("SpaceCenter", "Control_set_SAS", createArguments(arg))

	p, e := sc.conn.sendMessage(pr)
	res := &Response{}
	proto.Unmarshal(p, res)

	return
}

func (sc *SpaceCenter) ActivateNextStage(vessel []byte) (e error) {

	arg := [][]byte{vessel}
	pr := createRequest("SpaceCenter", "Control_ActivateNextStage", createArguments(arg))

	p, e := sc.conn.sendMessage(pr)
	if e != nil {
		return
	}
	res := &Response{}
	proto.Unmarshal(p, res)

	return
}

func (sc *SpaceCenter) GetThrottle() (val float32, e error) {
	arg := [][]byte{sc.Control}
	pr := createRequest("SpaceCenter", "Control_get_Throttle", createArguments(arg))
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
	arg := [][]byte{sc.Control, float32toByte(val)}
	pr := createRequest("SpaceCenter", "Control_set_Throttle", createArguments(arg))
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

func (sc *SpaceCenter) GetActionGroup(control []byte, group []byte) (e error) {
	//// TODO:
	return
}

func (sc *SpaceCenter) SetActionGroup(control []byte, group []byte, state bool) (e error) {
	//// TODO:
	return
}

func (sc *SpaceCenter) ToggleActionGroup(control []byte, group []byte) (e error) {
	//// TODO:
	return
}

// REQUEST AND ARGMUENT CREATION

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

func createArguments(args [][]byte) (arg []*Argument) {
	for pos, val := range args {
		arg = append(arg, &Argument{
			Position: uint32(pos),
			Value:    val,
		})
	}
	return
}
