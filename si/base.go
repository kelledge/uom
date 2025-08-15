package si

import (
	"github.com/kelledge/uom"
	"github.com/kelledge/uom/dim"
)

// SI base units (coherent)
var (
	Meter        = uom.DefineUnit[dim.Length]("m", 1.0)
	Second       = uom.DefineUnit[dim.Time]("s", 1.0)
	Kilogram     = uom.DefineUnit[dim.Mass]("kg", 1.0)
	Ampere       = uom.DefineUnit[dim.Current]("A", 1.0)
	Kelvin       = uom.DefineUnit[dim.Temperature]("K", 1.0)
	Mole         = uom.DefineUnit[dim.Substance]("mol", 1.0)
	Candela      = uom.DefineUnit[dim.Luminosity]("cd", 1.0)
	Scalar       = uom.DefineUnit[dim.Dimensionless]("scalar", 1.0)
	MeterSquared = uom.DefineUnit[dim.Area]("m^2", 1.0)
)
