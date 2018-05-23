package xorshift

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRand(t *testing.T) {
	r := rand.New(NewXorShift())
	r.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		fmt.Println(r.Float64())
	}
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(r.Int())
	// }
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(r.Intn(100))
	// }
}
