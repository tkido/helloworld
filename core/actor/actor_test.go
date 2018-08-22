package actor

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestInfection(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	a := Actor{"アン", 1}
	for i := 0; i < 20; i++ {
		a.Infect()
		fmt.Println(a)
	}
}
