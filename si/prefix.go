package si

import (
	"github.com/kelledge/uom"
	"github.com/kelledge/uom/dim"
)

// SI prefixes as dimensionless units
var (
	Yocto = uom.DefineUnit[dim.Dimensionless]("y", 1e-24)
	Zepto = uom.DefineUnit[dim.Dimensionless]("z", 1e-21)
	Atto  = uom.DefineUnit[dim.Dimensionless]("a", 1e-18)
	Femto = uom.DefineUnit[dim.Dimensionless]("f", 1e-15)
	Pico  = uom.DefineUnit[dim.Dimensionless]("p", 1e-12)
	Nano  = uom.DefineUnit[dim.Dimensionless]("n", 1e-9)
	Micro = uom.DefineUnit[dim.Dimensionless]("μ", 1e-6)
	Milli = uom.DefineUnit[dim.Dimensionless]("m", 1e-3)
	Centi = uom.DefineUnit[dim.Dimensionless]("c", 1e-2)
	Deci  = uom.DefineUnit[dim.Dimensionless]("d", 1e-1)
	Deca  = uom.DefineUnit[dim.Dimensionless]("da", 1e1)
	Hecto = uom.DefineUnit[dim.Dimensionless]("h", 1e2)
	Kilo  = uom.DefineUnit[dim.Dimensionless]("k", 1e3)
	Mega  = uom.DefineUnit[dim.Dimensionless]("M", 1e6)
	Giga  = uom.DefineUnit[dim.Dimensionless]("G", 1e9)
	Tera  = uom.DefineUnit[dim.Dimensionless]("T", 1e12)
	Peta  = uom.DefineUnit[dim.Dimensionless]("P", 1e15)
	Exa   = uom.DefineUnit[dim.Dimensionless]("E", 1e18)
	Zetta = uom.DefineUnit[dim.Dimensionless]("Z", 1e21)
	Yotta = uom.DefineUnit[dim.Dimensionless]("Y", 1e24)
)
