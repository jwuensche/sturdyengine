package sturdyengine_test

import (
	"os"
	"testing"

	"github.com/jwuensche/sturdy-engine"
	"github.com/op/go-logging"
)

func InitLogger(loggerName string) {
	backend1 := logging.NewLogBackend(os.Stderr, "", 0)
	backend1Formatter := logging.NewBackendFormatter(backend1, format)
	backend1Leveled := logging.AddModuleLevel(backend1)
	backend1Leveled.SetLevel(logging.ERROR, "")
	logging.SetBackend(backend1Leveled, backend1Formatter)
}

var log = logging.MustGetLogger("SturdyLog")
var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc}|%{shortfile} : %{level:.6s} %{id:03x}%{color:reset} %{message}`,
)
var backend = logging.NewLogBackend(os.Stderr, "", 0)

func TestInitalizeApi(t *testing.T) {
	c := sturdyengine.Connection{}
	if err := c.InitializeAPI("SturdyEngineTest"); err != nil {
		log.Error(err)
		t.FailNow()
	}
}
