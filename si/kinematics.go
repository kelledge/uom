package si

import (
	"github.com/kelledge/uom"
	"github.com/kelledge/uom/dim"
)

// Velocity
var (
	MeterPerSecond = uom.DefineUnit[dim.Velocity]("m/s", 1)
)

// Acceleration
var (
	MeterPerSecondSquared = uom.DefineUnit[dim.Acceleration]("m/s^2", 1)
)
