package conv

import "math"

//Littile Endian
//Convert float to bit representaion
//take less significant byte and fill it with lowest address.
//bit shift pattern eight places to the right and fill the next byte and so on.
func Float32ToBytes(fl float32, b []byte, j int) {
	f := math.Float32bits(fl)
	b[j] = byte(f)
	b[j+1] = byte(f >> 8)
	b[j+2] = byte(f >> 16)
	b[j+3] = byte(f >> 24)
}

func Float32To16intBytes(fl float32, b []byte, j int) {
	i := int16(32767 * fl)
	b[j] = byte(i)
	b[j+1] = byte(i >> 8)
}

func BytesToFloat32(b []byte, j int) float32 {
	var f uint32
	f |= uint32(b[j+3])
	f = f << 8
	f |= uint32(b[j+2])
	f = f << 8
	f |= uint32(b[j+1])
	f = f << 8
	f |= uint32(b[j])
	return math.Float32frombits(f)
}

func BytesToUint32(b []byte, j int) uint32 {
	var ui uint32
	ui |= uint32(b[j+3])
	ui = ui << 8
	ui |= uint32(b[j+2])
	ui = ui << 8
	ui |= uint32(b[j+1])
	ui = ui << 8
	ui |= uint32(b[j])
	return ui
}

func UInt32ToBytes(ui32 uint32, b []byte, j int) {
	b[j] = byte(ui32)
	b[j+1] = byte(ui32 >> 8)
	b[j+2] = byte(ui32 >> 16)
	b[j+3] = byte(ui32 >> 24)
}

func BytesToUint16(b []byte, j int) uint16 {
	var ui uint16
	ui |= uint16(b[j+1])
	ui = ui << 8
	ui |= uint16(b[j])
	return ui
}

func BytesToint16(b []byte, j int) int16 {
	var i int16
	i |= int16(b[j+1])
	i = i << 8
	i |= int16(b[j])
	return i
}

func Int16ToBytes(val int16, b []byte, i int) {
	b[i] = byte(val)
	b[i+1] = byte(val >> 8)
}
