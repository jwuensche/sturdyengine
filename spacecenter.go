package sturdyengine

import (
	"errors"
	"fmt"

	"github.com/golang/protobuf/proto"
	util "github.com/jwuensche/sturdyengine/internal"
	krpc "github.com/jwuensche/sturdyengine/internal/krpcproto"
)

// SpaceCenter saves all information used by `SpaceCenter` service of kRPC
type SpaceCenter struct {
	conn    *Connection
	Vessel  []byte
	Control []byte
}

//NewSpaceCenter creates a SpaceCenter instance with Vessel and Control being set to the current active Values
func NewSpaceCenter(conn *Connection) (sc SpaceCenter, e error) {
	if conn.Conn == nil {
		e = errors.New("Connection ist nil")
		return
	}
	sc = SpaceCenter{}
	sc.conn = conn
	return
}

//GetGameMode returns the current active GameMode (currently not implemented)
func (sc *SpaceCenter) GetGameMode() (r []byte, e error) {
	pr := createRequest("SpaceCenter", "get_GameMode", nil)

	p, e := sc.conn.sendMessage(pr)
	res := &krpc.Response{}
	proto.Unmarshal(p, res)

	//TODO: Translate to enums
	return
}

//SetPhysicsWarpFactor sets the physical warp factor on a range of 0 ... 3, this resets to 0 if onrailswarp is active
func (sc *SpaceCenter) SetPhysicsWarpFactor(factor uint8) (e error) {
	//krpc requires in this case the a multiplied value, this is only a fix for this, I do not know if this is intended behaviour
	arg := [][]byte{{2 * factor}}
	pr := createRequest("SpaceCenter", "set_PhysicsWarpFactor", createArguments(arg))
	_, e = sc.conn.sendMessage(pr)
	return
}

//GetPhysicsWarpFactor returns the current on a range of 0 ... 3
func (sc *SpaceCenter) GetPhysicsWarpFactor() (fac uint8, e error) {
	pr := createRequest("SpaceCenter", "get_PhysicsWarpFactor", nil)
	p, e := sc.conn.sendMessage(pr)
	res := &krpc.Response{}
	proto.Unmarshal(p, res)
	fac = res.GetResults()[0].GetValue()[0] / 2
	return
}

func (sc *SpaceCenter) SetWarp(factor uint64) (e error) {
	arg := [][]byte{util.Uint64ToByte(factor)}
	pr := createRequest("SpaceCenter", "set_RailsWarpFactor", createArguments(arg))
	_, e = sc.conn.sendMessage(pr)
	return
}

func (sc *SpaceCenter) GetWarp() (fac uint64, e error) {
	pr := createRequest("SpaceCenter", "get_RailsWarpFactor", nil)
	p, e := sc.conn.sendMessage(pr)
	res := &krpc.Response{}
	proto.Unmarshal(p, res)
	fac = util.ByteToUint64(res.GetResults()[0].GetValue())
	return
}

//GetMaximumWarpFactor returns the maximal available RailsWarpFactor at the moment
func (sc *SpaceCenter) GetMaximumWarpFactor() (fac uint64, e error) {
	pr := createRequest("SpaceCenter", "get_MaximumRailsWarpFactor", nil)
	p, e := sc.conn.sendMessage(pr)
	res := &krpc.Response{}
	proto.Unmarshal(p, res)
	fac = util.ByteToUint64(res.GetResults()[0].GetValue())
	return
}

func (sc *SpaceCenter) CheckWarpFactor(factor uint64) (result bool, e error) {
	arg := [][]byte{util.Uint64ToByte(factor)}
	pr := createRequest("SpaceCenter", "CanRailsWarpAt", createArguments(arg))
	p, e := sc.conn.sendMessage(pr)
	res := &krpc.Response{}
	proto.Unmarshal(p, res)
	fmt.Println(res.GetResults())
	result = util.ByteToBool(res.GetResults()[0].GetValue())
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
