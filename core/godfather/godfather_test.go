package godfather

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(len(data))
	max := len(data[0])
	for i := 0; i < 10; i++ {
		fmt.Println(data[0][rand.Intn(max)])
	}
}
