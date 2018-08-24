package sturdyengine_test

import (
	"fmt"
	"testing"

	"github.com/jwuensche/sturdyengine"
)

var scc, _ = sturdyengine.NewDefaultConnection()
var sc, e = sturdyengine.NewSpaceCenter(&scc)

func TestInitSpaceCenter(t *testing.T) {
	sc, e = sturdyengine.NewSpaceCenter(&scc)
	if e != nil {
		fmt.Println(e)
		t.FailNow()
	}
}
