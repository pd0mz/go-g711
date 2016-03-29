package g711

// MLawEncode encodes a slice of PCM16 samples to μ-law.
func MLawEncode(in []int16) []uint8 {
	if in == nil {
		return nil
	}
	out := make([]uint8, len(in))
	for i, s := range in {
		out[i] = MLawEncodeSample(s)
	}
	return out
}

// MLawDecode decodes a slice of μ-law samples to PCM16.
func MLawDecode(in []uint8) []int16 {
	if in == nil {
		return nil
	}
	out := make([]int16, len(in))
	for i, s := range in {
		out[i] = MLawDecodeSample(s)
	}
	return out
}

// MLawEncodeSample encodes a PCM16 sample to μ-law.
func MLawEncodeSample(s int16) uint8 {
	if s >= 0 {
		return μLawCompressTable[s>>4]
	}
	return 0x7f & μLawCompressTable[-s>>4]
}

// MLawDecodeSample decodes an μ-law sample to PCM16.
func MLawDecodeSample(s uint8) int16 {
	return μLawDecompressTable[s]
}
