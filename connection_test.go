package sturdyengine_test

import (
	"os"
	"testing"

	"github.com/jwuensche/sturdyengine"
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

//Declaring the connection on a global level to avoid redefining existing flag addre

var c = sturdyengine.Connection{}

func TestInitalizeApiAndClose(t *testing.T) {

	if err := c.InitializeAPI("SturdyEngineTest"); err != nil {
		log.Error(err)
		t.FailNow()
	}

}

func TestStatus(t *testing.T) {
	r, e := c.GetStatus()
	if e != nil {
		log.Error(e)
		t.FailNow()
	}
	log.Notice(r.String())
}
