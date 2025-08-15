package si

import (
	"github.com/kelledge/uom"
	"github.com/kelledge/uom/dim"
)

var (
	Newton            = uom.DefineUnit[dim.Force]("N", 1)
	NewtonMeter       = uom.DefineUnit[dim.Torque]("N·m", 1)
	Momentum          = uom.DefineUnit[dim.Momentum]("kg·m/s", 1)
	NewtonSecond      = uom.DefineUnit[dim.Impulse]("N·s", 1)
	NewtonMeterSecond = uom.DefineUnit[dim.AngularMomentum]("N·m·s", 1)
)

var (
	NewtonWeight = uom.DefineUnit[dim.Weight]("N", 1)
)
