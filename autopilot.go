package sturdyengine

//AutoPilot represents an instance of krpc autopilot for a vessel
type AutoPilot struct {
	sc  *SpaceCenter
	vsl *Vessel

	autopilot []byte
}

// AUTOPILOT - SPACECENTER.AUTOPILOT || SPACECENTER.VESSEL.AUTOPILOT

func (vsl *Vessel) NewAutoPilot() (ap AutoPilot, err error) {

	return
}

// EngageAutoPilot activates the given autopilot on the vessel it has been generated
func (sc *SpaceCenter) EngageAutoPilot(autopilot []byte) (e error) {

	return
}
