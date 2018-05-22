package xorshift

import (
	"fmt"
	"testing"
)

func TestXor(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(xor())
	}
}

var x uint64 = 88172645463325252

func xor() uint64 {
	x = x ^ (x << 7)
	x = x ^ (x >> 9)
	return x
}
