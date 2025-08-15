package si

import (
	"github.com/kelledge/uom"
	"github.com/kelledge/uom/dim"
)

// Mass units commonly used alongside kilogram base
var (
	// 1 g = 1e-3 kg
	Gram      = uom.DeriveUnit[dim.Mass]("g", uom.U(Milli), uom.U(Kilogram))
	Milligram = uom.DeriveUnit[dim.Mass]("mg", uom.U(Milli), uom.U(Gram))
	Microgram = uom.DeriveUnit[dim.Mass]("µg", uom.U(Micro), uom.U(Gram))

	// 1 tonne (t) = 1e3 kg — SI accepted unit
	Tonne = uom.DeriveUnit[dim.Mass]("t", uom.U(Kilo), uom.U(Kilogram))
)
