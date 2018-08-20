package sturdyengine_test

import (
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
	krpc "github.com/jwuensche/sturdyengine/internal/krpcproto"
	"github.com/jwuensche/sturdyengine/pkg"
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
	ser := krpc.Service{
		Name:          "Foo",
		Documentation: "Bar",
	}

	sers := krpc.Services{
		Services: []*krpc.Service{&ser},
	}

	p, e := proto.Marshal(&sers)
	if e != nil {
		log.Fatal(e)
		t.FailNow()
	}

	test := &krpc.Services{}

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
	log.Info("RPCs: " + strconv.FormatUint(r.GetRpcsExecuted(), 10))
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

	f, _ := os.Create("docs/procedure_doc.md")
	defer f.Close()

	for _, service := range r.GetServices() {
		f.WriteString("## " + service.GetName() + "\n")
		f.WriteString("#### Classes\n")
		for _, class := range service.GetClasses() {
			f.WriteString("```\n" + class.GetName() + "\n" + class.GetDocumentation() + "\n```\n")
		}
		for _, procedure := range service.GetProcedures() {
			f.WriteString("### " + procedure.GetName() + "\n\n")
			f.WriteString("```\n" + procedure.GetDocumentation() + "\n```\n")
			f.WriteString("#### Parameters \n")
			for _, parameter := range procedure.GetParameters() {
				f.WriteString("```\n")
				f.WriteString(parameter.GetName() + "		" + parameter.GetType().GetName() + "\n")
				f.WriteString("```\n")
			}
			f.WriteString("#### Return \n")
			f.WriteString("```\n" + procedure.GetReturnType().GetName() + "```\n")
		}
	}
}

func TestSpaceCenter(t *testing.T) {
	if c.Conn == nil {
		t.FailNow()
		return
	}
	//InitSpaceCenter
	sc, e := sturdyengine.InitSpaceCenter(&c)
	if e != nil {
		log.Info(e)
		t.FailNow()
	}
	//Quicksave
	e = sc.Quicksave()
	if e != nil {
		log.Info(e)
		t.FailNow()
	}
	//Get current vessel
	vessel, e := sc.GetActiveVessel()
	if e != nil {
		log.Info(e)
		t.FailNow()
	}
	//Get current vessels control
	control, e := sc.GetVesselControl(sc.Vessel)
	if e != nil {
		log.Info(e)
		t.FailNow()
	}
	//Test Control
	e = sc.ActivateNextStage(control)
	if e != nil {
		log.Info(e)
		t.FailNow()
	}
	// c.GetGameMode()
	e = sc.SetSAS(control, true)
	if e != nil {
		log.Info(e)
		t.FailNow()
	}
	sc.SetPhysicsWarpFactor(1)
	val, e := sc.GetPhysicsWarpFactor()
	if e != nil {
		log.Info(e)
		t.FailNow()
	}
	log.Info(val)
	sc.GetSAS(control)
	//Throttle
	sc.SetThrottle(0.54)
	_, e = sc.GetThrottle()
	if e != nil {
		log.Info(e)
		t.FailNow()
	}

	time.Sleep(2 * time.Second)
	e = sc.SetSAS(control, false)
	if e != nil {
		log.Info(e)
		t.FailNow()
	}

	sc.GetSAS(control)
	orb, e := sc.GetVesselOrbit(vessel)
	if e != nil {
		log.Info(e)
		t.FailNow()
	}
	alt, e := sc.GetApoapsisAltitude(orb)
	if e != nil {
		log.Info(e)
		t.FailNow()
	}
	log.Info(alt)
	time, e := sc.GetTimeToApoapsis(orb)
	if e != nil {
		log.Info(e)
		t.FailNow()
	}
	log.Info(time)
	alt, e = sc.GetRadius(orb)
	if e != nil {
		log.Info(e)
		t.FailNow()
	}
	//Test this further not sure about it yet
	spd, e := sc.GetSpeed(orb)
	log.Info(spd)
	if e != nil {
		log.Info(e)
		t.FailNow()
	}
	log.Info(alt)
	//Quickload
	e = sc.Quickload()
	if e != nil {
		log.Info(e)
		t.FailNow()
	}
}

func TestClose(t *testing.T) {
	if c.Conn == nil {
		t.FailNow()
		return
	}
	c.Close()
}
