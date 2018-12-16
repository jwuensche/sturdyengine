package sturdyengine

import (
	"github.com/golang/protobuf/proto"
	util "github.com/jwuensche/sturdyengine/internal"
	krpc "github.com/jwuensche/sturdyengine/internal/krpcproto"
)

//Vessel represents a combination of krpc Vessel & Control allowing direct operations on the current vessel
type Vessel struct {
	sc *SpaceCenter

	vessel  []byte
	control []byte
}

//NewVessel returns a initialized instance of a vessel ready to use
func (sc *SpaceCenter) NewVessel() (vessel Vessel, err error) {
	vsl, err := sc.GetActiveVessel()
	vessel = Vessel{
		sc:     sc,
		vessel: vsl,
	}
	vessel.control, err = vessel.GetVesselControl()
	return
}

//ActivateNextStage activate the next stage of the given control unit
func (vsl *Vessel) ActivateNextStage() (err error) {
	arg := [][]byte{vsl.control}
	pr := createRequest("SpaceCenter", "Control_ActivateNextStage", createArguments(arg))

	p, err := vsl.sc.conn.sendMessage(pr)
	if err != nil {
		return
	}
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

//GetVesselControl returns control object to a given vessel
func (vsl *Vessel) GetVesselControl() (r []byte, e error) {
	arg := [][]byte{vsl.vessel}
	pr := createRequest("SpaceCenter", "Vessel_get_Control", createArguments(arg))

	p, e := vsl.sc.conn.sendMessage(pr)
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
func (vsl *Vessel) GetSAS() (r bool, e error) {
	arg := [][]byte{vsl.control}
	pr := createRequest("SpaceCenter", "Control_get_SAS", createArguments(arg))
	p, e := vsl.sc.conn.sendMessage(pr)
	res := &krpc.Response{}
	proto.Unmarshal(p, res)
	r = util.ByteToBool(res.GetResults()[0].GetValue())
	return
}

//SetSAS sets the state of the given control unit's SAS, true to activate
func (vsl *Vessel) SetSAS(state bool) (e error) {
	s := util.BoolToByte(state)
	arg := [][]byte{vsl.vessel, s}
	pr := createRequest("SpaceCenter", "Control_set_SAS", createArguments(arg))

	_, e = vsl.sc.conn.sendMessage(pr)
	return
}

//GetThrottle returns the Throtlle of the given control unit as float value between 0 and 1
func (vsl *Vessel) GetThrottle() (val float32, e error) {
	arg := [][]byte{vsl.control}
	pr := createRequest("SpaceCenter", "Control_get_Throttle", createArguments(arg))
	p, e := vsl.sc.conn.sendMessage(pr)
	if e != nil {
		return
	}
	res := &krpc.Response{}
	proto.Unmarshal(p, res)
	val = util.ByteToFloat32(res.GetResults()[0].GetValue())
	return
}

//SetThrottle sets the Throtlle of the given control unit as float value between 0 and 1
func (vsl *Vessel) SetThrottle(val float32) (e error) {
	arg := [][]byte{vsl.control, util.Float32toByte(val)}
	pr := createRequest("SpaceCenter", "Control_set_Throttle", createArguments(arg))
	_, e = vsl.sc.conn.sendMessage(pr)

	return
}

func (vsl *Vessel) GetActionGroup(group []byte) (e error) {
	//// TODO:
	return
}

func (vsl *Vessel) SetActionGroup(group []byte, state bool) (e error) {
	//// TODO:
	return
}

func (vsl *Vessel) ToggleActionGroup(group []byte) (e error) {
	//// TODO:
	return
}
