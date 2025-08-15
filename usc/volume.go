package usc

import (
	"github.com/kelledge/uom"
	"github.com/kelledge/uom/dim"
)

// Length-cubed volumes
var (
	CubicInch = uom.DeriveUnit[dim.Volume]("in^3", uom.U(Inch).Pow(3))
	CubicFoot = uom.DeriveUnit[dim.Volume]("ft^3", uom.U(Foot).Pow(3))
	CubicYard = uom.DeriveUnit[dim.Volume]("yd^3", uom.U(Yard).Pow(3))
)

// Liquid measures
var (
	Gallon     = uom.DeriveUnit[dim.Volume]("gal", uom.U(Inch).Pow(3).MulScalar(231))
	Quart      = uom.DeriveUnit[dim.Volume]("qt", uom.U(Gallon).DivScalar(4))
	Pint       = uom.DeriveUnit[dim.Volume]("pt", uom.U(Gallon).DivScalar(8))
	Cup        = uom.DeriveUnit[dim.Volume]("cup", uom.U(Gallon).DivScalar(16))
	FluidOunce = uom.DeriveUnit[dim.Volume]("fl oz", uom.U(Gallon).DivScalar(128))
	Tablespoon = uom.DeriveUnit[dim.Volume]("tbsp", uom.U(FluidOunce).DivScalar(2))
	Teaspoon   = uom.DeriveUnit[dim.Volume]("tsp", uom.U(Tablespoon).DivScalar(3))

	// Petroleum barrel (42 gallons exactly)
	Barrel = uom.DeriveUnit[dim.Volume]("bbl", uom.U(Gallon).MulScalar(42))
)
