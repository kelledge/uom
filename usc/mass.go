package usc

import (
	"github.com/kelledge/uom"
	"github.com/kelledge/uom/dim"
)

var (
	PoundMass = uom.DefineUnit[dim.Mass]("lb", 0.45359237)
	OunceMass = uom.DefineUnit[dim.Mass]("oz", 0.028349523125)
	ShortTon  = uom.DefineUnit[dim.Mass]("ton", 907.18474) // US short ton
)
