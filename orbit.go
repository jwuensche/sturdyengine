package sturdyengine

import (
	"github.com/golang/protobuf/proto"
	util "github.com/jwuensche/sturdyengine/internal"
	krpc "github.com/jwuensche/sturdyengine/internal/krpcproto"
)

//Orbit represents a snapshot of a vessel's orbit
type Orbit struct {
	sc    *SpaceCenter
	orbit []byte
}

// ORBIT - SPACECENTER.VESSEL.ORBIT || SPACECENTER.CELESTIALBODY.ORBIT

//GetVesselOrbit returns a snapshot of the current orbit the vessel is on
func (vsl *Vessel) GetVesselOrbit() (orbit Orbit, e error) {
	arg := [][]byte{vsl.vessel}
	orbit = Orbit{
		sc: vsl.sc,
	}
	pr := createRequest("SpaceCenter", "Vessel_get_Orbit", createArguments(arg))
	p, e := orbit.sc.conn.sendMessage(pr)
	res := &krpc.Response{}
	proto.Unmarshal(p, res)
	orbit.orbit = res.GetResults()[0].GetValue()
	return
}

//GetApoapsisAltitude returns the apoapsis of the orbit relative to the surface the reference object in meters
func (orb *Orbit) GetApoapsisAltitude() (alt float64, e error) {
	alt, e = orb.getOrbitInfo("Orbit_get_ApoapsisAltitude")
	return
}

//GetPeriapsisAltitude returns the periapsis of the orbit relative to the surface the reference object in meters
func (orb *Orbit) GetPeriapsisAltitude() (alt float64, e error) {
	alt, e = orb.getOrbitInfo("Orbit_get_PeriapsisAltitude")
	return
}

//GetSemiMajorAxis returns the semi-major axis in meters
func (orb *Orbit) GetSemiMajorAxis() (alt float64, e error) {
	alt, e = orb.getOrbitInfo("Orbit_get_SemiMajorAxis")
	return
}

//GetSemiMinorAxis returns the semi-minor axis in meters
func (orb *Orbit) GetSemiMinorAxis() (alt float64, e error) {
	alt, e = orb.getOrbitInfo("Orbit_get_SemiMinorAxis")
	return
}

//GetRadius returns the current radius to the center of mass of the reference object in meters
func (orb *Orbit) GetRadius() (alt float64, e error) {
	alt, e = orb.getOrbitInfo("Orbit_get_Radius")
	return
}

//GetSpeed returns the current speed of the object in meters
func (orb *Orbit) GetSpeed() (alt float64, e error) {
	alt, e = orb.getOrbitInfo("Orbit_get_Speed")
	return
}

//GetPeriod returns the orbital period in seconds
func (orb *Orbit) GetPeriod() (time float64, e error) {
	time, e = orb.getOrbitInfo("Orbit_get_Period")
	return
}

//GetTimeToApoapsis returns the time to apoapsis in seconds
func (orb *Orbit) GetTimeToApoapsis() (time float64, e error) {
	time, e = orb.getOrbitInfo("Orbit_get_TimeToApoapsis")
	return
}

//GetTimeToApoapsis returns the time to periapsis in seconds
func (orb *Orbit) GetTimeToPeriapsis() (time float64, e error) {
	time, e = orb.getOrbitInfo("Orbit_get_TimeToPeriapsis")
	return
}

//GetEpoch returns the time since the epoch at which mean anomaly at epoch was measured in seconds
func (orb *Orbit) GetEpoch() (time float64, e error) {
	time, e = orb.getOrbitInfo("Orbit_get_Epoch")
	return
}

// GetEccentricity returns the eccentricity of the given orbit
func (orb *Orbit) GetEccentricity() (ecc float32, e error) {
	ecc, e = orb.getOrbitInfoFloat32("Orbit_get_Eccentricity")
	return
}

func (orb *Orbit) getOrbitInfo(procedure string) (alt float64, e error) {
	arg := [][]byte{orb.orbit}
	pr := createRequest("SpaceCenter", procedure, createArguments(arg))
	p, e := orb.sc.conn.sendMessage(pr)
	if e != nil {
		return
	}
	res := &krpc.Response{}
	proto.Unmarshal(p, res)
	alt = util.ByteToFloat64(res.GetResults()[0].GetValue())
	return
}

func (orb *Orbit) getOrbitInfoFloat32(procedure string) (alt float32, e error) {
	arg := [][]byte{orb.orbit}
	pr := createRequest("SpaceCenter", procedure, createArguments(arg))
	p, e := orb.sc.conn.sendMessage(pr)
	if e != nil {
		return
	}
	res := &krpc.Response{}
	proto.Unmarshal(p, res)
	alt = util.ByteToFloat32(res.GetResults()[0].GetValue())
	return
}
