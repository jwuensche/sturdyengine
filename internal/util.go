package util

import (
	"encoding/binary"
	"math"
)

func ByteToFloat32(r []byte) (f32 float32) {
	bits := binary.LittleEndian.Uint32(r)
	f32 = math.Float32frombits(bits)
	return
}

func Float32toByte(r float32) (b []byte) {
	bits := math.Float32bits(r)
	b = make([]byte, 4)
	binary.LittleEndian.PutUint32(b, bits)
	return
}

func BoolToByte(state bool) (b []byte) {
	if state {
		b = []byte{byte(1)}
	} else {
		b = []byte{byte(0)}
	}
	return
}

func ByteToFloat64(r []byte) (f64 float64) {
	bits := binary.LittleEndian.Uint64(r)
	f64 = math.Float64frombits(bits)
	return
}

func ByteToBool(r []byte) (b bool) {
	if r[0] == 1 {
		b = true
	}
	return
}

func Uint64ToByte(r uint64) (b []byte) {
	b = make([]byte, 8)
	binary.LittleEndian.PutUint64(b, r)
	return
}

func ByteToUint64(r []byte) (u64 uint64) {
	u64 = binary.LittleEndian.Uint64(r)
	return
}
