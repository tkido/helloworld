package godfather

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestNewMaleSingle(t *testing.T) {
	fmt.Println("男性単体名")
	sin := data[MaleSingle]
	max := len(sin)
	for i := 0; i < 40; i++ {
		s := sin[rand.Intn(max)]
		fmt.Println(s)
	}
}

func TestNewMaleCompound(t *testing.T) {
	fmt.Println("男性合成名")
	pre := data[MalePrefix]
	suf := data[MaleSuffix]
	maxPre := len(pre)
	maxSuf := len(suf)
	for i := 0; i < 40; i++ {
		s := pre[rand.Intn(maxPre)] + suf[rand.Intn(maxSuf)]
		fmt.Println(s)
	}
}

func TestNewFemaleSingle(t *testing.T) {
	fmt.Println("女性単体名")
	sin := data[FemaleSingle]
	max := len(sin)
	for i := 0; i < 40; i++ {
		s := sin[rand.Intn(max)]
		fmt.Println(s)
	}
}

func TestNewFemaleCompound(t *testing.T) {
	fmt.Println("女性合成名")
	pre := data[FemalePrefix]
	suf := data[FemaleSuffix]
	maxPre := len(pre)
	maxSuf := len(suf)
	for i := 0; i < 40; i++ {
		s := pre[rand.Intn(maxPre)] + suf[rand.Intn(maxSuf)]
		fmt.Println(s)
	}
}
