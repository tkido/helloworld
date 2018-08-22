package actor

import (
	"fmt"
	"math/rand"
)

// Actor is actor
type Actor struct {
	Name      string
	Infection int
}

// Actor Type
const (
	Human = iota
	Vampire
	Zombie
)

func (a *Actor) String() string {
	return fmt.Sprintf("%s:%3d", a.Name, a.Infection)
}

// Infect is Infect
func (a *Actor) Infect() {
	ini := a.Infection
	for i := 0; i < ini; i++ {
		if a.Infection < rand.Intn(100)+1 {
			a.Infection++
		}
	}
}