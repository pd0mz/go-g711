package g711

// ALawEncode encodes a slice of PCM16 samples to a-law.
func ALawEncode(in []int16) []uint8 {
	out := make([]uint8, len(in))
	for i, s := range in {
		out[i] = ALawEncodeSample(s)
	}
	return out
}

// ALawDecode decodes a slice of a-law samples to PCM16.
func ALawDecode(in []uint8) []int16 {
	out := make([]int16, len(in))
	for i, s := range in {
		out[i] = ALawDecodeSample(s)
	}
	return out
}

// ALawEncodeSample encodes a PCM16 sample to a-law.
func ALawEncodeSample(s int16) uint8 {
	if s >= 0 {
		return aLawCompressTable[s>>3]
	}
	return 0x7f & aLawCompressTable[-s>>3]
}

// ALawDecodeSample decodes an a-law sample to PCM16.
func ALawDecodeSample(s uint8) int16 {
	return aLawDecompressTable[s]
}
