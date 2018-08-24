package sturdyengine_test

import (
	"fmt"
	"testing"

	"github.com/jwuensche/sturdyengine"
)

var c, _ = sturdyengine.NewDefaultConnection()

func TestStatus(t *testing.T) {
	_, e := c.GetStatus()
	if e != nil {
		fmt.Println(e)
		t.FailNow()
	}
}

func TestServices(t *testing.T) {
	_, e := c.GetServices()
	if e != nil {
		fmt.Println(e)
		t.FailNow()
	}
}
