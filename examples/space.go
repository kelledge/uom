package examples

import (
	"github.com/kelledge/uom"
	"github.com/kelledge/uom/dim"
	"github.com/kelledge/uom/std"
	"github.com/kelledge/uom/usc"
)

type Engine struct {
	Name        string
	Agency      string
	Propellants string
	ThrustSL    uom.Quantity[dim.Force]
	ThrustVac   uom.Quantity[dim.Force]
	IspSL       uom.Quantity[dim.Time]
	IspVac      uom.Quantity[dim.Time]
	ChamberP    uom.Quantity[dim.Pressure]
	Notes       string
}

var (
	EngineF1 = Engine{
		Name:        "Rocketdyne F-1",
		Agency:      "NASA/Rocketdyne (historic)",
		Propellants: "LOX/RP-1 (gas-generator)",
		ThrustSL:    usc.PoundsForce.Of(1_522_000),
		ThrustVac:   usc.PoundsForce.Of(1_748_200),
		IspSL:       std.Second.Of(265.4),
		IspVac:      std.Second.Of(304.1),
		ChamberP:    usc.PSI.Of(982.0),
		Notes:       "Saturn V S-IC; uprated post-Apollo 8.",
	}
)
