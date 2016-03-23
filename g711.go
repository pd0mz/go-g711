// Package g711 implements the ITU-T standard for audio companding.
package g711

const (
	SignBit   = 0x80 // Sign bit for an aLaw byte
	QuantMask = 0x0f // Quantization field mask
	Segments  = 8    // Number of aLaw segments
	SegShift  = 4    // Left shift for segment number
	SegMask   = 0x70 // Segment field mask

)

var (
	segEnd = []uint16{0xFF, 0x1FF, 0x3FF, 0x7FF, 0xFFF, 0x1FFF, 0x3FFF, 0x7FFF}
)

func search(val uint16, table []uint16, size uint16) uint16 {
	var i uint16
	for i = 0; i < size; i++ {
		if val <= table[i] {
			return i
		}
	}
	return size
}
