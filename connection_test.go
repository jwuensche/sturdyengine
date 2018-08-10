package sturdyengine_test

import (
	"os"
	"testing"

	"github.com/golang/protobuf/proto"
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

//Declaring the connection on a global level to avoid redefining existing flag address

var c = sturdyengine.Connection{}

func TestServiceProto(t *testing.T) {
	ser := sturdyengine.Service{
		Name:          "Foo",
		Documentation: "Bar",
	}

	sers := sturdyengine.Services{
		Services: []*sturdyengine.Service{&ser},
	}

	p, e := proto.Marshal(&sers)
	if e != nil {
		log.Fatal(e)
		t.FailNow()
	}

	test := &sturdyengine.Services{}

	proto.Unmarshal(p, test)

	log.Info("Raw: " + test.String())

	if sers.GetServices()[0].Name != test.GetServices()[0].Name {
		log.Fatal("Unmarshaled unequal to original structure")
		t.FailNow()
	}

}

func TestInitalizeApiAndClose(t *testing.T) {
	if err := c.InitializeAPI("SturdyEngineTest"); err != nil {
		log.Error(err)
		t.FailNow()
	}
}

func TestStatus(t *testing.T) {
	if c.Conn == nil {
		t.FailNow()
		return
	}
	r, e := c.GetStatus()
	if e != nil {
		log.Error(e)
		t.FailNow()
	}
	log.Info("Raw : " + r.String())
	log.Info("kRPC Server Version: " + r.GetVersion())
	log.Info(r.GetRpcsExecuted())
}

func TestServices(t *testing.T) {
	if c.Conn == nil {
		t.FailNow()
		return
	}
	r, e := c.GetServices()
	if e != nil {
		log.Error(e)
		t.FailNow()
	}

	for _, service := range r.GetServices() {
		log.Info(service.Name)
		log.Info("--------------")
		for _, procedure := range service.GetProcedures() {
			log.Info(procedure.GetName())
			log.Info(procedure.GetDocumentation())
		}
	}
}

func TestClose(t *testing.T) {
	if c.Conn == nil {
		t.FailNow()
		return
	}
	c.Close()
}
