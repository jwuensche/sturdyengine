package main

import (
	"os"
	"time"

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

	vessel, err := sc.NewVessel()
	if err != nil {
		os.Exit(1)
	}

	sc.Quicksave()

	vessel.SetThrottle(0.75)
	vessel.ActivateNextStage()
	time.Sleep(10 * time.Second)
	sc.Quickload()

}
