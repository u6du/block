package xor

func Copy(dst, src []byte) {
	for i := range dst {
		dst[i] ^= src[i]
	}
}
func Copy32(dst, src *[32]byte) {
	for i := range dst {
		dst[i] ^= src[i]
	}
}
