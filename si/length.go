package si

import (
	"github.com/kelledge/uom"
	"github.com/kelledge/uom/dim"
)

var (
	Kilometer  = uom.DeriveUnit[dim.Length]("km", uom.U(Kilo), uom.U(Meter))
	Centimeter = uom.DeriveUnit[dim.Length]("cm", uom.U(Centi), uom.U(Meter))
	Millimeter = uom.DeriveUnit[dim.Length]("mm", uom.U(Milli), uom.U(Meter))
	Micrometer = uom.DeriveUnit[dim.Length]("µm", uom.U(Micro), uom.U(Meter))
	Nanometer  = uom.DeriveUnit[dim.Length]("nm", uom.U(Nano), uom.U(Meter))
)
