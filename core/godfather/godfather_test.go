package godfather

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"bitbucket.org/tkido/gostock/my/json"
)

func TestNew(t *testing.T) {
	names := [][]string{}
	err := json.Load("./testdata/names.json", &names)
	if err != nil {
		t.Fatal(err)
	}
	rand.Seed(time.Now().UnixNano())
	fmt.Println(len(names))
	max := len(names[0])
	for i := 0; i < 10; i++ {
		fmt.Println(names[0][rand.Intn(max)])
	}
}
