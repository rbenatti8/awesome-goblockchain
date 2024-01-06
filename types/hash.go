package types

import (
	"crypto/rand"
)

type Hash [32]uint8

func (h *Hash) IsZero() bool {
	for _, b := range h {
		if b != 0 {
			return false
		}
	}
	return true
}
func (h *Hash) ToBytes() []byte {
	return h[:]
}

func HashFromBytes(b []byte) Hash {
	if len(b) != 32 {
		panic("HashFromBytes: invalid length")
	}

	var h Hash
	copy(h[:], b)
	return h
}

func randomBytes(size int) []byte {
	b := make([]byte, size)
	_, _ = rand.Read(b)

	return b
}

func RandomHash() Hash {
	return HashFromBytes(randomBytes(32))
}
