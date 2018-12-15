package main

import (
	"os"

	"github.com/jwuensche/sturdyengine"
)

func main() {
	conn, err := sturdyengine.NewDefaultConnection()
	if err != nil {
		os.Exit(1)
	}
	sc, err := sturdyengine.NewSpaceCenter(&conn)
	if err != nil {
		os.Exit(1)
	}
	sc.SetThrottle(1)
	sc.SetSAS(sc.Vessel, true)
	sc.ActivateNextStage(sc.Control)
}
