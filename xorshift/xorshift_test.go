package xorshift

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"io"
	"testing"
	"time"
)

const intSize = 32 << (^uint(0) >> 63)

func TestXor(t *testing.T) {
	Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		fmt.Println(Float64())
	}
	for i := 0; i < 10; i++ {
		fmt.Println(Int())
	}
}

func TestSeed(t *testing.T) {
	i := time.Now().UnixNano()
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(i))
	fmt.Printf("%v\n", i)
	fmt.Printf("%x\n", i)
	fmt.Printf("%x\n", b)
}
func TestMd5(t *testing.T) {
	fmt.Println(intSize)
	h := md5.New()

	io.WriteString(h, "testa")
	sum := h.Sum(nil)
	fmt.Printf("%x\n", sum)
	fmt.Println(len(sum))

	var u0, u1 uint64
	var i uint
	for i = 0; i < 8; i++ {
		fmt.Printf("%x\n", u0)
		fmt.Printf("%x\n", sum[i])
		u0 = u0<<8 | uint64(sum[i])
	}
	for i = 0; i < 8; i++ {
		u1 = u1<<8 | uint64(sum[i+8])
	}
	fmt.Printf("%x\n", u0)
	fmt.Printf("%x\n", u1)

}
