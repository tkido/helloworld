package xorshift

import (
	"math/rand"
)

const (
	max  = 1 << 63
	mask = max - 1
)

// NewXorShift returns rand source
func NewXorShift() rand.Source64 {
	return &XorShift{88172645463325252}
}

// XorShift is
type XorShift struct {
	i uint64
}

// Seed uses the provided seed value to initialize the generator to a deterministic state.
func (x *XorShift) Seed(seed int64) {
	x.i = uint64(seed)
	for j := 0; j < 64; j++ {
		x.Uint64()
	}
}

// Uint64 returns a non-negative pseudo-random 64-bit integer as an uint64.
func (x *XorShift) Uint64() uint64 {
	x.i = x.i ^ (x.i << 7)
	x.i = x.i ^ (x.i >> 9)
	return x.i
}

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
func (x *XorShift) Int63() int64 {
	return int64(x.Uint64() & mask)
}
