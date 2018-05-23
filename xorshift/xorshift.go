package xorshift

import (
	"math"
)

// default seed
var x uint64 = 88172645463325252
var max = float64(math.MaxUint64)

// Seed gave seed
func Seed(seed int64) {
	x = uint64(seed)
}

func next() uint64 {
	x = x ^ (x << 7)
	x = x ^ (x >> 9)
	return x
}

// Float64 returns, as a float64, a pseudo-random number in [0.0,1.0) from the default Source.
func Float64() float64 {
	return float64(next()) / max
}

// Int returns a non-negative pseudo-random int from the default Source.
func Int() int {
	return int(next())
}
