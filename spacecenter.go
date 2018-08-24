package sturdyengine

import (
	"github.com/golang/protobuf/proto"
	krpc "github.com/jwuensche/sturdyengine/internal/krpcproto"
)

//InitSpaceCenter creates a SpaceCenter instance with Vessel and Control being set to the current active Values
func InitSpaceCenter(conn *Connection) (sc SpaceCenter, e error) {
	sc = SpaceCenter{}
	sc.conn = conn
	sc.Vessel, e = sc.GetActiveVessel()
	sc.Control, e = sc.GetVesselControl(sc.Vessel)
	return
}

//GetGameMode returns the current active GameMode (currently not implemented)
func (sc *SpaceCenter) GetGameMode() (r []byte, e error) {
	pr := createRequest("SpaceCenter", "get_GameMode", nil)

	p, e := sc.conn.sendMessage(pr)
	res := &krpc.Response{}
	proto.Unmarshal(p, res)

	return
}

//GetActiveVessel returns the current active Vessel as byte slice
func (sc *SpaceCenter) GetActiveVessel() (r []byte, e error) {
	pr := createRequest("SpaceCenter", "get_ActiveVessel", nil)

	p, e := sc.conn.sendMessage(pr)
	res := &krpc.Response{}
	proto.Unmarshal(p, res)
	r = res.GetResults()[0].GetValue()
	return
}

//SetPhysicsWarpFactor sets the physical warp factor on a range of 0 ... 3, this resets to 0 if onrailswarp is active
func (sc *SpaceCenter) SetPhysicsWarpFactor(factor uint8) (e error) {
	arg := [][]byte{[]byte{2 * factor}}
	pr := createRequest("SpaceCenter", "set_PhysicsWarpFactor", createArguments(arg))
	_, e = sc.conn.sendMessage(pr)
	return
}

func (sc *SpaceCenter) GetPhysicsWarpFactor() (fac uint8, e error) {
	pr := createRequest("SpaceCenter", "get_PhysicsWarpFactor", nil)
	p, e := sc.conn.sendMessage(pr)
	res := &krpc.Response{}
	proto.Unmarshal(p, res)
	fac = res.GetResults()[0].GetValue()[0] / 2
	return
}

func (sc *SpaceCenter) SetWarp(factor uint64) (e error) {
	arg := [][]byte{uint64ToByte(factor)}
	pr := createRequest("SpaceCenter", "set_RailsWarpFactor", createArguments(arg))
	_, e = sc.conn.sendMessage(pr)
	return
}

func (sc *SpaceCenter) GetWarp() (fac uint64, e error) {
	pr := createRequest("SpaceCenter", "get_RailsWarpFactor", nil)
	p, e := sc.conn.sendMessage(pr)
	res := &krpc.Response{}
	proto.Unmarshal(p, res)
	fac = byteToUint64(res.GetResults()[0].GetValue())
	return
}

func (sc *SpaceCenter) GetMaximumWarpFactor() (fac uint64, e error) {
	pr := createRequest("SpaceCenter", "get_MaximumRailsWarpFactor", nil)
	p, e := sc.conn.sendMessage(pr)
	res := &krpc.Response{}
	proto.Unmarshal(p, res)
	fac = byteToUint64(res.GetResults()[0].GetValue())
	return
}

// ORBIT - SPACECENTER.VESSEL.ORBIT || SPACECENTER.CELESTIALBODY.ORBIT

//GetVesselOrbit returns a snapshot of the current orbit the vessel is on
func (sc *SpaceCenter) GetVesselOrbit(vessel []byte) (orb []byte, e error) {
	arg := [][]byte{vessel}
	pr := createRequest("SpaceCenter", "Vessel_get_Orbit", createArguments(arg))
	p, e := sc.conn.sendMessage(pr)
	res := &krpc.Response{}
	proto.Unmarshal(p, res)
	orb = res.GetResults()[0].GetValue()
	return
}

//GetApoapsisAltitude returns the apoapsis of the orbit relative to the surface the reference object in meters
func (sc *SpaceCenter) GetApoapsisAltitude(orbit []byte) (alt float64, e error) {
	alt, e = sc.getOrbitInfo(orbit, "Orbit_get_ApoapsisAltitude")
	return
}

//GetPeriapsisAltitude returns the periapsis of the orbit relative to the surface the reference object in meters
func (sc *SpaceCenter) GetPeriapsisAltitude(orbit []byte) (alt float64, e error) {
	alt, e = sc.getOrbitInfo(orbit, "Orbit_get_PeriapsisAltitude")
	return
}

//GetSemiMajorAxis returns the semi-major axis in meters
func (sc *SpaceCenter) GetSemiMajorAxis(orbit []byte) (alt float64, e error) {
	alt, e = sc.getOrbitInfo(orbit, "Orbit_get_SemiMajorAxis")
	return
}

//GetSemiMinorAxis returns the semi-minor axis in meters
func (sc *SpaceCenter) GetSemiMinorAxis(orbit []byte) (alt float64, e error) {
	alt, e = sc.getOrbitInfo(orbit, "Orbit_get_SemiMinorAxis")
	return
}

//GetRadius returns the current radius to the center of mass of the reference object in meters
func (sc *SpaceCenter) GetRadius(orbit []byte) (alt float64, e error) {
	alt, e = sc.getOrbitInfo(orbit, "Orbit_get_Radius")
	return
}

//GetSpeed returns the current speed of the object in meters
func (sc *SpaceCenter) GetSpeed(orbit []byte) (alt float64, e error) {
	alt, e = sc.getOrbitInfo(orbit, "Orbit_get_Speed")
	return
}

//GetPeriod returns the orbital period in seconds
func (sc *SpaceCenter) GetPeriod(orbit []byte) (time float64, e error) {
	time, e = sc.getOrbitInfo(orbit, "Orbit_get_Period")
	return
}

//GetTimeToApoapsis returns the time to apoapsis in seconds
func (sc *SpaceCenter) GetTimeToApoapsis(orbit []byte) (time float64, e error) {
	time, e = sc.getOrbitInfo(orbit, "Orbit_get_TimeToApoapsis")
	return
}

//GetTimeToApoapsis returns the time to periapsis in seconds
func (sc *SpaceCenter) GetTimeToPeriapsis(orbit []byte) (time float64, e error) {
	time, e = sc.getOrbitInfo(orbit, "Orbit_get_TimeToPeriapsis")
	return
}

//GetEpoch returns the time since the epoch at which mean anomaly at epoch was measured in seconds
func (sc *SpaceCenter) GetEpoch(orbit []byte) (time float64, e error) {
	time, e = sc.getOrbitInfo(orbit, "Orbit_get_Epoch")
	return
}

// GetEccentricity returns the eccentricity of the given orbit
func (sc *SpaceCenter) GetEccentricity(orbit []byte) (ecc float32, e error) {
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
	res := &krpc.Response{}
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
	res := &krpc.Response{}
	proto.Unmarshal(p, res)
	alt = byteToFloat32(res.GetResults()[0].GetValue())
	return
}

// CONTROL - SPACECENTER.VESSEL.CONTROL

//GetVesselControl returns control object to a given vessel
func (sc *SpaceCenter) GetVesselControl(vessel []byte) (r []byte, e error) {
	arg := [][]byte{vessel}
	pr := createRequest("SpaceCenter", "Vessel_get_Control", createArguments(arg))

	p, e := sc.conn.sendMessage(pr)
	if e != nil {
		return
	}
	res := &krpc.Response{}
	e = proto.Unmarshal(p, res)
	if e != nil {
		return
	}

	r = res.GetResults()[0].GetValue()
	return
}

//GetSAS returns the current state of the given control unit SAS, true in case it is activated
func (sc *SpaceCenter) GetSAS(control []byte) (r bool, e error) {
	arg := [][]byte{control}
	pr := createRequest("SpaceCenter", "Control_get_SAS", createArguments(arg))
	p, e := sc.conn.sendMessage(pr)
	res := &krpc.Response{}
	proto.Unmarshal(p, res)
	r = byteToBool(res.GetResults()[0].GetValue())
	return
}

//SetSAS sets the state of the given control unit's SAS, true to activate
func (sc *SpaceCenter) SetSAS(vessel []byte, state bool) (e error) {
	s := boolToByte(state)
	arg := [][]byte{vessel, s}
	pr := createRequest("SpaceCenter", "Control_set_SAS", createArguments(arg))

	_, e = sc.conn.sendMessage(pr)
	return
}

//ActivateNextStage activate the next stage of the given control unit
func (sc *SpaceCenter) ActivateNextStage(control []byte) (e error) {

	arg := [][]byte{control}
	pr := createRequest("SpaceCenter", "Control_ActivateNextStage", createArguments(arg))

	p, e := sc.conn.sendMessage(pr)
	if e != nil {
		return
	}
	res := &krpc.Response{}
	proto.Unmarshal(p, res)

	return
}

//GetThrottle returns the Throtlle of the given control unit as float value between 0 and 1
func (sc *SpaceCenter) GetThrottle() (val float32, e error) {
	arg := [][]byte{sc.Control}
	pr := createRequest("SpaceCenter", "Control_get_Throttle", createArguments(arg))
	p, e := sc.conn.sendMessage(pr)
	if e != nil {
		return
	}
	res := &krpc.Response{}
	proto.Unmarshal(p, res)
	val = byteToFloat32(res.GetResults()[0].GetValue())
	return
}

//SetThrottle sets the Throtlle of the given control unit as float value between 0 and 1
func (sc *SpaceCenter) SetThrottle(val float32) (e error) {
	arg := [][]byte{sc.Control, float32toByte(val)}
	pr := createRequest("SpaceCenter", "Control_set_Throttle", createArguments(arg))
	_, e = sc.conn.sendMessage(pr)

	return
}

//Quicksave creates a Quicksave
func (sc *SpaceCenter) Quicksave() (e error) {
	pr := createRequest("SpaceCenter", "Quicksave", nil)
	_, e = sc.conn.sendMessage(pr)
	return
}

//Quickload loads the last used Quicksave
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

func createRequest(service string, procedure string, arguments []*krpc.Argument) (pr *krpc.Request) {
	pc := &krpc.ProcedureCall{
		Service:   service,
		Procedure: procedure,
		Arguments: arguments,
	}
	pr = &krpc.Request{
		Calls: []*krpc.ProcedureCall{pc},
	}
	return
}

func createArguments(args [][]byte) (arg []*krpc.Argument) {
	for pos, val := range args {
		arg = append(arg, &krpc.Argument{
			Position: uint32(pos),
			Value:    val,
		})
	}
	return
}
