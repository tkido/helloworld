package godfather

import (
	"math/rand"
	"time"
)

// for data.go
const (
	MaleSingle = iota
	MalePrefix
	MaleSuffix
	FemaleSingle
	FemalePrefix
	FemaleSuffix
)

func init() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		rand.Int()
	}
}

// Next take next name
func Next(gender int) string {
	max := len(data[0])
	return data[0][rand.Intn(max)]
}
