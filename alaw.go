package g711

import (
	"encoding/binary"
	"errors"
	"io"
)

// ALawEncode encodes a slice of PCM16 samples to a-law.
func ALawEncode(in []int16) []uint8 {
	if in == nil {
		return nil
	}
	out := make([]uint8, len(in))
	for i, s := range in {
		out[i] = ALawEncodeSample(s)
	}
	return out
}

// ALawDecode decodes a slice of a-law samples to PCM16.
func ALawDecode(in []uint8) []int16 {
	if in == nil {
		return nil
	}
	out := make([]int16, len(in))
	for i, s := range in {
		out[i] = ALawDecodeSample(s)
	}
	return out
}

// ALawEncodeSample encodes a PCM16 sample to a-law.
func ALawEncodeSample(s int16) uint8 {
	if s >= 0 {
		return aLawCompressTable[s>>4]
	}
	return 0x7f & aLawCompressTable[-s>>4]
}

// ALawDecodeSample decodes an a-law sample to PCM16.
func ALawDecodeSample(s uint8) int16 {
	return aLawDecompressTable[s]
}

type ALawEncoder struct {
	io.Reader
	r io.Reader
}

// NewALawEncoder builds an io.Writer that consumes an io.Reader.
func NewALawEncoder(r io.Reader) (*ALawEncoder, error) {
	if r == nil {
		return nil, errNilReader
	}
	return &ALawEncoder{r: r}, nil
}

func (e *ALawEncoder) Read(p []byte) (n int, err error) {
	if p == nil {
		return 0, errNilBuffer
	}

	var (
		l = len(p)
		b = make([]byte, l<<1)
	)
	if _, err2 := e.r.Read(b); err2 != nil {
		return 0, err2
	}
	for i := 0; i < l; i++ {
		s := int16(binary.BigEndian.Uint16(b[i<<1:]))
		p[i] = ALawEncodeSample(s)
	}
	return len(p), nil
}

type ALawDecoder struct {
	io.Reader
	binary.ByteOrder
	r io.Reader
}

// NewALawDecoder builds an io.Writer that consumes an io.Reader.
func NewALawDecoder(r io.Reader) (*ALawDecoder, error) {
	if r == nil {
		return nil, errors.New("g711: nil reader")
	}
	return &ALawDecoder{
		r:         r,
		ByteOrder: binary.BigEndian,
	}, nil
}

func (e *ALawDecoder) Read(p []byte) (n int, err error) {
	var (
		l = len(p)
		h = l >> 1
		b = make([]byte, h)
	)
	if _, err2 := e.r.Read(b); err2 != nil {
		return 0, err2
	}
	for i := 0; i < h; i++ {
		s := ALawDecodeSample(b[i])
		e.ByteOrder.PutUint16(p[i<<1:], uint16(s))
	}
	return len(p), nil
}

// Interface check
var _ io.Reader = (*ALawEncoder)(nil)
var _ io.Reader = (*ALawDecoder)(nil)
