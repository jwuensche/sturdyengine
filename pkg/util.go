package sturdyengine

import (
	"encoding/binary"
	"math"
)

func byteToFloat32(r []byte) (f32 float32) {
	bits := binary.LittleEndian.Uint32(r)
	f32 = math.Float32frombits(bits)
	return
}

func float32toByte(r float32) (b []byte) {
	bits := math.Float32bits(r)
	b = make([]byte, 4)
	binary.LittleEndian.PutUint32(b, bits)
	return
}

func boolTobyte(state bool) (b []byte) {
	if state {
		b = []byte{byte(1)}
	} else {
		b = []byte{byte(0)}
	}
	return
}
