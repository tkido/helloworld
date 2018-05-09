package godfather

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Next take next name
func Next() string {
	max := len(data[0])
	return data[0][rand.Intn(max)]
}
